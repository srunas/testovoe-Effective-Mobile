package subscription_repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/srunas/effective-mobile/internal/infrastructure/repository/subscription-repository/sqlcgen"
)

type Implementation struct {
	q *sqlcgen.Queries
}

func NewImplementation(db *sqlx.DB) *Implementation {
	return &Implementation{q: sqlcgen.New(db)}
}
