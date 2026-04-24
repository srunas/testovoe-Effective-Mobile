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

func TestGetSubscription_Success(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	id := subscription.ID("sub-123")
	sub := &subscription.Subscription{
		ID:          id,
		ServiceName: "Spotify",
		Price:       299,
		UserID:      "user-1",
		StartDate:   time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
	}

	repo.On("GetByID", context.Background(), id).Return(sub, nil)

	result, err := svc.GetSubscription(context.Background(), service.GetSubscriptionRequest{ID: string(id)})

	assert.NoError(t, err)
	assert.Equal(t, sub, result.Subscription)
	repo.AssertExpectations(t)
}

func TestGetSubscription_NotFound(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	id := subscription.ID("missing")
	repo.On("GetByID", context.Background(), id).Return(nil, nil)

	result, err := svc.GetSubscription(context.Background(), service.GetSubscriptionRequest{ID: string(id)})

	assert.ErrorIs(t, err, service.ErrNotFound)
	assert.Nil(t, result.Subscription)
	repo.AssertExpectations(t)
}

func TestGetSubscription_RepoError(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	id := subscription.ID("sub-123")
	repoErr := errors.New("db error")
	repo.On("GetByID", context.Background(), id).Return(nil, repoErr)

	_, err := svc.GetSubscription(context.Background(), service.GetSubscriptionRequest{ID: string(id)})

	assert.ErrorIs(t, err, repoErr)
	repo.AssertExpectations(t)
}
