package subscription_service

import (
	"context"
	"log/slog"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
	"github.com/srunas/effective-mobile/internal/domain/service"
)

func (s *Implementation) GetSubscription(
	ctx context.Context,
	req service.GetSubscriptionRequest,
) (service.GetSubscriptionResponse, error) {
	sub, err := s.repo.GetByID(ctx, subscription.ID(req.ID))
	if err != nil {
		slog.ErrorContext(ctx, "ошибка получения подписки", "id", req.ID, "error", err)
		return service.GetSubscriptionResponse{}, err
	}

	if sub == nil {
		return service.GetSubscriptionResponse{}, service.ErrNotFound
	}

	return service.GetSubscriptionResponse{Subscription: sub}, nil
}
