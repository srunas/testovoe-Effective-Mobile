package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/srunas/effective-mobile/internal/domain/service"
)

// @Summary      Удалить подписку
// @Description  Удаляет подписку по ID
// @Tags         subscriptions
// @Param        id   path      string  true  "ID подписки"
// @Success      204
// @Failure      404  {string}  string  "подписка не найдена"
// @Failure      500  {string}  string  "внутренняя ошибка сервера"
// @Router       /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id обязателен", http.StatusBadRequest)
		return
	}

	if err := h.svc.DeleteSubscription(r.Context(), service.DeleteSubscriptionRequest{ID: id}); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, "подписка не найдена", http.StatusNotFound)
			return
		}

		slog.ErrorContext(r.Context(), "ошибка удаления подписки", "id", id, "error", err)
		http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
