package service

import (
	"context"
	"errors"
	"time"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
)

var ErrNotFound = errors.New("подписка не найдена")

type CreateSubscriptionRequest struct {
	ServiceName string
	Price       int64
	UserID      string
	StartDate   time.Time
	EndDate     *time.Time
}

type CreateSubscriptionResponse struct {
	ID string
}

type UpdateSubscriptionRequest struct {
	ID          string
	ServiceName string
	Price       int64
	StartDate   time.Time
	EndDate     *time.Time
}

type DeleteSubscriptionRequest struct {
	ID string
}

type GetSubscriptionRequest struct {
	ID string
}

type GetSubscriptionResponse struct {
	Subscription *subscription.Subscription
}

type ListSubscriptionsRequest struct {
	UserID string
}

type ListSubscriptionsResponse struct {
	Subscriptions []*subscription.Subscription
}

type CalculateTotalRequest struct {
	UserID      string
	ServiceName string
	From        time.Time
	To          time.Time
}

type CalculateTotalResponse struct {
	Total int64
}

type Subscription interface {
	CreateSubscription(ctx context.Context, req CreateSubscriptionRequest) (CreateSubscriptionResponse, error)
	UpdateSubscription(ctx context.Context, req UpdateSubscriptionRequest) error
	DeleteSubscription(ctx context.Context, req DeleteSubscriptionRequest) error
	GetSubscription(ctx context.Context, req GetSubscriptionRequest) (GetSubscriptionResponse, error)
	ListSubscriptions(ctx context.Context, req ListSubscriptionsRequest) (ListSubscriptionsResponse, error)
	CalculateTotal(ctx context.Context, req CalculateTotalRequest) (CalculateTotalResponse, error)
}
