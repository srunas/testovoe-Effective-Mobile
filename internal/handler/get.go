package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/srunas/effective-mobile/internal/domain/service"
)

// @Summary      Получить подписку
// @Description  Возвращает подписку по ID
// @Tags         subscriptions
// @Produce      json
// @Param        id   path      string  true  "ID подписки"
// @Success      200  {object}  subscription.Subscription
// @Failure      404  {string}  string  "подписка не найдена"
// @Failure      500  {string}  string  "внутренняя ошибка сервера"
// @Router       /subscriptions/{id} [get]
func (h *SubscriptionHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id обязателен", http.StatusBadRequest)
		return
	}

	result, err := h.svc.GetSubscription(r.Context(), service.GetSubscriptionRequest{ID: id})
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, "подписка не найдена", http.StatusNotFound)
			return
		}

		slog.ErrorContext(r.Context(), "ошибка получения подписки", "id", id, "error", err)
		http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if encodeErr := json.NewEncoder(w).Encode(result.Subscription); encodeErr != nil {
		slog.ErrorContext(r.Context(), "ошибка сериализации ответа", "error", encodeErr)
	}
}
