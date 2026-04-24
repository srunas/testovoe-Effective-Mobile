package subscription_service

import (
	"context"
	"log/slog"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
	"github.com/srunas/effective-mobile/internal/domain/service"
)

func (s *Implementation) CreateSubscription(
	ctx context.Context,
	req service.CreateSubscriptionRequest,
) (service.CreateSubscriptionResponse, error) {
	sub := subscription.New(req.ServiceName, req.Price, req.UserID, req.StartDate, req.EndDate)

	if err := s.repo.Create(ctx, sub); err != nil {
		slog.ErrorContext(ctx, "ошибка создания подписки", "error", err)
		return service.CreateSubscriptionResponse{}, err
	}

	slog.InfoContext(ctx, "подписка создана", "id", string(sub.ID))

	return service.CreateSubscriptionResponse{ID: string(sub.ID)}, nil
}
