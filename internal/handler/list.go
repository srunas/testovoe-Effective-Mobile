package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/srunas/effective-mobile/internal/domain/service"
)

// @Summary      Список подписок
// @Description  Возвращает список подписок, опционально фильтруя по user_id
// @Tags         subscriptions
// @Produce      json
// @Param        user_id  query     string  false  "ID пользователя"
// @Success      200      {array}   subscription.Subscription
// @Failure      500      {string}  string  "внутренняя ошибка сервера"
// @Router       /subscriptions [get]
func (h *SubscriptionHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	result, err := h.svc.ListSubscriptions(r.Context(), service.ListSubscriptionsRequest{UserID: userID})
	if err != nil {
		slog.ErrorContext(r.Context(), "ошибка получения списка подписок", "error", err)
		http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if encodeErr := json.NewEncoder(w).Encode(result.Subscriptions); encodeErr != nil {
		slog.ErrorContext(r.Context(), "ошибка сериализации ответа", "error", encodeErr)
	}
}
