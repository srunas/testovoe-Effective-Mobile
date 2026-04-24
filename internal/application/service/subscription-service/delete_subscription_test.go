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

func TestDeleteSubscription_Success(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	id := subscription.ID("sub-123")
	existing := &subscription.Subscription{
		ID:          id,
		ServiceName: "Netflix",
		Price:       999,
		UserID:      "user-1",
		StartDate:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	repo.On("GetByID", context.Background(), id).Return(existing, nil)
	repo.On("Delete", context.Background(), id).Return(nil)

	err := svc.DeleteSubscription(context.Background(), service.DeleteSubscriptionRequest{ID: string(id)})

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteSubscription_NotFound(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	id := subscription.ID("missing")
	repo.On("GetByID", context.Background(), id).Return(nil, nil)

	err := svc.DeleteSubscription(context.Background(), service.DeleteSubscriptionRequest{ID: string(id)})

	assert.ErrorIs(t, err, service.ErrNotFound)
	repo.AssertExpectations(t)
}

func TestDeleteSubscription_RepoDeleteError(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	id := subscription.ID("sub-123")
	existing := &subscription.Subscription{
		ID:          id,
		ServiceName: "Netflix",
		Price:       999,
		UserID:      "user-1",
		StartDate:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	repoErr := errors.New("db error")

	repo.On("GetByID", context.Background(), id).Return(existing, nil)
	repo.On("Delete", context.Background(), id).Return(repoErr)

	err := svc.DeleteSubscription(context.Background(), service.DeleteSubscriptionRequest{ID: string(id)})

	assert.ErrorIs(t, err, repoErr)
	repo.AssertExpectations(t)
}
