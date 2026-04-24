package subscription_service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	subscriptionservice "github.com/srunas/effective-mobile/internal/application/service/subscription-service"
	"github.com/srunas/effective-mobile/internal/domain/service"
)

func TestCreateSubscription_Success(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	req := service.CreateSubscriptionRequest{
		ServiceName: "Netflix",
		Price:       999,
		UserID:      "user-1",
		StartDate:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	repo.On("Create", mock.Anything, mock.Anything).Return(nil)

	result, err := svc.CreateSubscription(context.Background(), req)

	assert.NoError(t, err)
	assert.NotEmpty(t, result.ID)
	repo.AssertExpectations(t)
}

func TestCreateSubscription_RepoError(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	req := service.CreateSubscriptionRequest{
		ServiceName: "Netflix",
		Price:       999,
		UserID:      "user-1",
		StartDate:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	repoErr := errors.New("db error")
	repo.On("Create", mock.Anything, mock.Anything).Return(repoErr)

	result, err := svc.CreateSubscription(context.Background(), req)

	assert.ErrorIs(t, err, repoErr)
	assert.Empty(t, result.ID)
	repo.AssertExpectations(t)
}
