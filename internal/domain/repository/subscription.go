package repository

import (
	"context"
	"time"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
)

type Subscription interface {
	Create(ctx context.Context, s *subscription.Subscription) error
	GetByID(ctx context.Context, id subscription.ID) (*subscription.Subscription, error)
	Update(ctx context.Context, s *subscription.Subscription) error
	Delete(ctx context.Context, id subscription.ID) error
	List(ctx context.Context, userID string) ([]*subscription.Subscription, error)
	CalculateTotal(ctx context.Context, userID string, serviceName string, from time.Time, to time.Time) (int64, error)
}
