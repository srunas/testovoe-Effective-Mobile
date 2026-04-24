package handler

import (
	"github.com/go-chi/chi/v5"

	"github.com/srunas/effective-mobile/internal/domain/service"
)

type SubscriptionHandler struct {
	svc service.Subscription
}

func NewSubscriptionHandler(svc service.Subscription) *SubscriptionHandler {
	return &SubscriptionHandler{svc: svc}
}

func (h *SubscriptionHandler) Register(r chi.Router) {
	r.Route("/subscriptions", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Get("/", h.List)
		r.Get("/total", h.CalculateTotal)
		r.Get("/{id}", h.Get)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
	})
}
