package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/srunas/effective-mobile/internal/domain/service"
)

// @Summary      Подсчёт суммы подписок
// @Description  Возвращает суммарную стоимость подписок за период. Фильтрация по user_id и service_name — опциональна
// @Tags         subscriptions
// @Produce      json
// @Param        from          query     string  true   "Начало периода (MM-YYYY)"
// @Param        to            query     string  true   "Конец периода (MM-YYYY)"
// @Param        user_id       query     string  false  "ID пользователя"
// @Param        service_name  query     string  false  "Название сервиса"
// @Success      200           {object}  map[string]int64
// @Failure      400           {string}  string  "некорректные параметры"
// @Failure      500           {string}  string  "внутренняя ошибка сервера"
// @Router       /subscriptions/total [get]
func (h *SubscriptionHandler) CalculateTotal(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	serviceName := r.URL.Query().Get("service_name")
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	if fromStr == "" || toStr == "" {
		http.Error(w, "параметры from и to обязательны (формат MM-YYYY)", http.StatusBadRequest)
		return
	}

	from, err := time.Parse("01-2006", fromStr)
	if err != nil {
		http.Error(w, "некорректный формат from, ожидается MM-YYYY", http.StatusBadRequest)
		return
	}

	to, err := time.Parse("01-2006", toStr)
	if err != nil {
		http.Error(w, "некорректный формат to, ожидается MM-YYYY", http.StatusBadRequest)
		return
	}

	result, err := h.svc.CalculateTotal(r.Context(), service.CalculateTotalRequest{
		UserID:      userID,
		ServiceName: serviceName,
		From:        from,
		To:          to,
	})
	if err != nil {
		slog.ErrorContext(r.Context(), "ошибка подсчёта суммы", "error", err)
		http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if encodeErr := json.NewEncoder(w).Encode(map[string]int64{"total": result.Total}); encodeErr != nil {
		slog.ErrorContext(r.Context(), "ошибка сериализации ответа", "error", encodeErr)
	}
}
