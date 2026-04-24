package subscription_repository

import (
	"context"
	"time"

	"github.com/srunas/effective-mobile/internal/infrastructure/repository/subscription-repository/sqlcgen"
)

func (r *Implementation) CalculateTotal(
	ctx context.Context, userID string, serviceName string, from time.Time, to time.Time,
) (int64, error) {
	total, err := r.q.CalculateTotal(ctx, sqlcgen.CalculateTotalParams{
		Column1:     userID,
		Column2:     serviceName,
		StartDate:   from,
		StartDate_2: to,
	})
	if err != nil {
		return 0, err
	}

	return total, nil
}
