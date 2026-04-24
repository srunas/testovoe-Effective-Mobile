package subscription_service

import (
	"context"
	"log/slog"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
	"github.com/srunas/effective-mobile/internal/domain/service"
)

func (s *Implementation) UpdateSubscription(
	ctx context.Context,
	req service.UpdateSubscriptionRequest,
) error {
	sub, err := s.repo.GetByID(ctx, subscription.ID(req.ID))
	if err != nil {
		return err
	}

	if sub == nil {
		return service.ErrNotFound
	}

	sub.ServiceName = req.ServiceName
	sub.Price = req.Price
	sub.StartDate = req.StartDate
	sub.EndDate = req.EndDate

	if err = s.repo.Update(ctx, sub); err != nil {
		slog.ErrorContext(ctx, "ошибка обновления подписки", "id", req.ID, "error", err)
		return err
	}

	slog.InfoContext(ctx, "подписка обновлена", "id", req.ID)

	return nil
}
