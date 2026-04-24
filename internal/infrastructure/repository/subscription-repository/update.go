package subscription_repository

import (
	"context"
	"database/sql"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
	"github.com/srunas/effective-mobile/internal/infrastructure/repository/subscription-repository/sqlcgen"
)

func (r *Implementation) Update(ctx context.Context, s *subscription.Subscription) error {
	params := sqlcgen.UpdateSubscriptionParams{
		ID:          string(s.ID),
		ServiceName: s.ServiceName,
		Price:       s.Price,
		StartDate:   s.StartDate,
	}

	if s.EndDate != nil {
		params.EndDate = sql.NullTime{Time: *s.EndDate, Valid: true}
	}

	return r.q.UpdateSubscription(ctx, params)
}
