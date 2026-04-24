package subscription_repository

import (
	"context"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
)

func (r *Implementation) List(ctx context.Context, userID string) ([]*subscription.Subscription, error) {
	rows, err := r.q.ListSubscriptions(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]*subscription.Subscription, 0, len(rows))
	for _, row := range rows {
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
		result = append(result, s)
	}

	return result, nil
}
