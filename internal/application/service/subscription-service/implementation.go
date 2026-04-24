package subscription_service

import "github.com/srunas/effective-mobile/internal/domain/repository"

type Implementation struct {
	repo repository.Subscription
}

func NewImplementation(repo repository.Subscription) *Implementation {
	return &Implementation{repo: repo}
}
