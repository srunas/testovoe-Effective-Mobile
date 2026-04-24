package subscription_repository

import (
	"context"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
)

func (r *Implementation) Delete(ctx context.Context, id subscription.ID) error {
	return r.q.DeleteSubscription(ctx, string(id))
}
