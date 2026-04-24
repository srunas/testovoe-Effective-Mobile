package subscription_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
)

func (r *Implementation) GetByID(ctx context.Context, id subscription.ID) (*subscription.Subscription, error) {
	row, err := r.q.GetSubscriptionByID(ctx, string(id))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil //nolint:nilnil // nil без ошибки означает "не найдено"
	}
	if err != nil {
		return nil, err
	}

	s := &subscription.Subscription{
		ID:          subscription.ID(row.ID),
		ServiceName: row.ServiceName,
		Price:       row.Price,
		UserID:      row.UserID,
		StartDate:   row.StartDate,
	}

	if row.EndDate.Valid {
		s.EndDate = &row.EndDate.Time
	}

	return s, nil
}
