package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/srunas/effective-mobile/internal/domain/service"
)

type updateRequest struct {
	ServiceName string  `json:"service_name"`
	Price       int64   `json:"price"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
}

// @Summary      Обновить подписку
// @Description  Обновляет данные существующей подписки
// @Tags         subscriptions
// @Accept       json
// @Param        id    path      string        true  "ID подписки"
// @Param        body  body      updateRequest true  "Новые данные подписки"
// @Success      204
// @Failure      400  {string}  string  "некорректный запрос"
// @Failure      404  {string}  string  "подписка не найдена"
// @Failure      500  {string}  string  "внутренняя ошибка сервера"
// @Router       /subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id обязателен", http.StatusBadRequest)
		return
	}

	var req updateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "некорректный запрос", http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse("01-2006", req.StartDate)
	if err != nil {
		http.Error(w, "некорректный формат start_date, ожидается MM-YYYY", http.StatusBadRequest)
		return
	}

	svcReq := service.UpdateSubscriptionRequest{
		ID:          id,
		ServiceName: req.ServiceName,
		Price:       req.Price,
		StartDate:   startDate,
	}

	if req.EndDate != nil {
		endDate, parseErr := time.Parse("01-2006", *req.EndDate)
		if parseErr != nil {
			http.Error(w, "некорректный формат end_date, ожидается MM-YYYY", http.StatusBadRequest)
			return
		}

		svcReq.EndDate = &endDate
	}

	if err = h.svc.UpdateSubscription(r.Context(), svcReq); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, "подписка не найдена", http.StatusNotFound)
			return
		}

		slog.ErrorContext(r.Context(), "ошибка обновления подписки", "id", id, "error", err)
		http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
