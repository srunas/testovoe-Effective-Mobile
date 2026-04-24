package subscription_service

import (
	"context"
	"log/slog"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
	"github.com/srunas/effective-mobile/internal/domain/service"
)

func (s *Implementation) DeleteSubscription(
	ctx context.Context,
	req service.DeleteSubscriptionRequest,
) error {
	sub, err := s.repo.GetByID(ctx, subscription.ID(req.ID))
	if err != nil {
		return err
	}

	if sub == nil {
		return service.ErrNotFound
	}

	if err = s.repo.Delete(ctx, subscription.ID(req.ID)); err != nil {
		slog.ErrorContext(ctx, "ошибка удаления подписки", "id", req.ID, "error", err)
		return err
	}

	slog.InfoContext(ctx, "подписка удалена", "id", req.ID)

	return nil
}
