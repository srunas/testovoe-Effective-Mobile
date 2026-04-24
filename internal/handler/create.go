package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/srunas/effective-mobile/internal/domain/service"
)

type createRequest struct {
	ServiceName string  `json:"service_name"`
	Price       int64   `json:"price"`
	UserID      string  `json:"user_id"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
}

// @Summary      Создать подписку
// @Description  Создаёт новую запись о подписке пользователя
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        body  body      createRequest  true  "Данные подписки"
// @Success      201   {object}  map[string]string
// @Failure      400   {string}  string  "некорректный запрос"
// @Failure      500   {string}  string  "внутренняя ошибка сервера"
// @Router       /subscriptions [post]
func (h *SubscriptionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "некорректный запрос", http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse("01-2006", req.StartDate)
	if err != nil {
		http.Error(w, "некорректный формат start_date, ожидается MM-YYYY", http.StatusBadRequest)
		return
	}

	svcReq := service.CreateSubscriptionRequest{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
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

	result, err := h.svc.CreateSubscription(r.Context(), svcReq)
	if err != nil {
		slog.ErrorContext(r.Context(), "ошибка создания подписки", "error", err)
		http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(map[string]string{"id": result.ID}); err != nil {
		slog.ErrorContext(r.Context(), "ошибка сериализации ответа", "error", err)
	}
}
