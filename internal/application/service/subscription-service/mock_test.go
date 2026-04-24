package subscription_service_test

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) Create(ctx context.Context, s *subscription.Subscription) error {
	args := m.Called(ctx, s)
	return args.Error(0)
}

func (m *mockRepo) GetByID(ctx context.Context, id subscription.ID) (*subscription.Subscription, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*subscription.Subscription), args.Error(1)
}

func (m *mockRepo) Update(ctx context.Context, s *subscription.Subscription) error {
	args := m.Called(ctx, s)
	return args.Error(0)
}

func (m *mockRepo) Delete(ctx context.Context, id subscription.ID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockRepo) List(ctx context.Context, userID string) ([]*subscription.Subscription, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*subscription.Subscription), args.Error(1)
}

func (m *mockRepo) CalculateTotal(
	ctx context.Context,
	userID string,
	serviceName string,
	from time.Time,
	to time.Time,
) (int64, error) {
	args := m.Called(ctx, userID, serviceName, from, to)
	return args.Get(0).(int64), args.Error(1)
}
