package subscription_service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	subscriptionservice "github.com/srunas/effective-mobile/internal/application/service/subscription-service"
	"github.com/srunas/effective-mobile/internal/domain/entity/subscription"
	"github.com/srunas/effective-mobile/internal/domain/service"
)

func TestUpdateSubscription_Success(t *testing.T) {
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

	req := service.UpdateSubscriptionRequest{
		ID:          string(id),
		ServiceName: "Netflix Premium",
		Price:       1299,
		StartDate:   time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
	}

	repo.On("GetByID", context.Background(), id).Return(existing, nil)
	repo.On("Update", context.Background(), mock.Anything).Return(nil)

	err := svc.UpdateSubscription(context.Background(), req)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUpdateSubscription_NotFound(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	id := subscription.ID("missing")
	repo.On("GetByID", context.Background(), id).Return(nil, nil)

	err := svc.UpdateSubscription(context.Background(), service.UpdateSubscriptionRequest{ID: string(id)})

	assert.ErrorIs(t, err, service.ErrNotFound)
	repo.AssertExpectations(t)
}

func TestUpdateSubscription_RepoGetError(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	id := subscription.ID("sub-123")
	repoErr := errors.New("db error")
	repo.On("GetByID", context.Background(), id).Return(nil, repoErr)

	err := svc.UpdateSubscription(context.Background(), service.UpdateSubscriptionRequest{ID: string(id)})

	assert.ErrorIs(t, err, repoErr)
	repo.AssertExpectations(t)
}
