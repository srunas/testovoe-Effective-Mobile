package subscription_service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	subscriptionservice "github.com/srunas/effective-mobile/internal/application/service/subscription-service"
	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
	"github.com/srunas/effective-mobile/internal/domain/service"
)

func TestListSubscriptions_Success(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	subs := []*subscription.Subscription{
		{
			ID:          "sub-1",
			ServiceName: "Netflix",
			Price:       999,
			UserID:      "user-1",
			StartDate:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:          "sub-2",
			ServiceName: "Spotify",
			Price:       299,
			UserID:      "user-1",
			StartDate:   time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	repo.On("List", context.Background(), "user-1").Return(subs, nil)

	result, err := svc.ListSubscriptions(context.Background(), service.ListSubscriptionsRequest{UserID: "user-1"})

	assert.NoError(t, err)
	assert.Len(t, result.Subscriptions, 2)
	repo.AssertExpectations(t)
}

func TestListSubscriptions_Empty(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	repo.On("List", context.Background(), "").Return([]*subscription.Subscription{}, nil)

	result, err := svc.ListSubscriptions(context.Background(), service.ListSubscriptionsRequest{})

	assert.NoError(t, err)
	assert.Empty(t, result.Subscriptions)
	repo.AssertExpectations(t)
}

func TestListSubscriptions_RepoError(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	repoErr := errors.New("db error")
	repo.On("List", context.Background(), "user-1").Return(nil, repoErr)

	_, err := svc.ListSubscriptions(context.Background(), service.ListSubscriptionsRequest{UserID: "user-1"})

	assert.ErrorIs(t, err, repoErr)
	repo.AssertExpectations(t)
}
