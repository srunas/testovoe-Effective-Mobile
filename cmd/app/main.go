// @title           Effective Mobile Subscription API
// @version         1.0
// @description     REST сервис для агрегации данных об онлайн подписках
// @host            localhost:8080
// @BasePath        /

package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/srunas/effective-mobile/docs"
	subscriptionservice "github.com/srunas/effective-mobile/internal/application/service/subscription-service"
	"github.com/srunas/effective-mobile/internal/config"
	"github.com/srunas/effective-mobile/internal/handler"
	subscriptionrepository "github.com/srunas/effective-mobile/internal/infrastructure/repository/subscription-repository"
)

const shutdownTimeout = 5 * time.Second

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Get()

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		slog.Error("ошибка подключения к БД", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("подключение к БД установлено")

	repo := subscriptionrepository.NewImplementation(db)
	svc := subscriptionservice.NewImplementation(repo)
	subscriptionHandler := handler.NewSubscriptionHandler(svc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	subscriptionHandler.Register(r)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second, //nolint:mnd // стандартный таймаут заголовка
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		slog.Info("сервер запущен", "port", cfg.Server.Port)
		if listenErr := srv.ListenAndServe(); listenErr != nil && !errors.Is(listenErr, http.ErrServerClosed) {
			slog.Error("ошибка сервера", "error", listenErr)
		}
	}()

	<-quit

	slog.Info("остановка сервера")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if shutdownErr := srv.Shutdown(ctx); shutdownErr != nil {
		slog.Error("ошибка при остановке сервера", "error", shutdownErr)
	}
}
