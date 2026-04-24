package subscription_service

import (
	"context"
	"log/slog"

	"github.com/srunas/effective-mobile/internal/domain/service"
)

func (s *Implementation) ListSubscriptions(
	ctx context.Context,
	req service.ListSubscriptionsRequest,
) (service.ListSubscriptionsResponse, error) {
	subscriptions, err := s.repo.List(ctx, req.UserID)
	if err != nil {
		slog.ErrorContext(ctx, "ошибка получения списка подписок", "error", err)
		return service.ListSubscriptionsResponse{}, err
	}

	return service.ListSubscriptionsResponse{Subscriptions: subscriptions}, nil
}
