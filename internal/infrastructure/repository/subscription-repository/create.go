package subscription_repository

import (
	"context"
	"database/sql"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
	"github.com/srunas/effective-mobile/internal/infrastructure/repository/subscription-repository/sqlcgen"
)

func (r *Implementation) Create(ctx context.Context, s *subscription.Subscription) error {
	params := sqlcgen.CreateSubscriptionParams{
		ID:          string(s.ID),
		ServiceName: s.ServiceName,
		Price:       s.Price,
		UserID:      s.UserID,
		StartDate:   s.StartDate,
	}

	if s.EndDate != nil {
		params.EndDate = sql.NullTime{Time: *s.EndDate, Valid: true}
	}

	return r.q.CreateSubscription(ctx, params)
}
