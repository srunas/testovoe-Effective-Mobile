package subscription_service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	subscriptionservice "github.com/srunas/effective-mobile/internal/application/service/subscription-service"
	"github.com/srunas/effective-mobile/internal/domain/service"
)

func TestCalculateTotal_Success(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	from := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)

	repo.On("CalculateTotal", context.Background(), "user-1", "Netflix", from, to).Return(int64(1998), nil)

	result, err := svc.CalculateTotal(context.Background(), service.CalculateTotalRequest{
		UserID:      "user-1",
		ServiceName: "Netflix",
		From:        from,
		To:          to,
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(1998), result.Total)
	repo.AssertExpectations(t)
}

func TestCalculateTotal_NoFilters(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	from := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)

	repo.On("CalculateTotal", context.Background(), "", "", from, to).Return(int64(5000), nil)

	result, err := svc.CalculateTotal(context.Background(), service.CalculateTotalRequest{
		From: from,
		To:   to,
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(5000), result.Total)
	repo.AssertExpectations(t)
}

func TestCalculateTotal_RepoError(t *testing.T) {
	repo := &mockRepo{}
	svc := subscriptionservice.NewImplementation(repo)

	from := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	repoErr := errors.New("db error")

	repo.On("CalculateTotal", context.Background(), "", "", from, to).Return(int64(0), repoErr)

	_, err := svc.CalculateTotal(context.Background(), service.CalculateTotalRequest{From: from, To: to})

	assert.ErrorIs(t, err, repoErr)
	repo.AssertExpectations(t)
}
