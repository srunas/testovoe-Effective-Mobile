package subscription

import (
	"time"

	"github.com/google/uuid"
)

type ID string

type Subscription struct {
	ID          ID         `json:"id"`
	ServiceName string     `json:"service_name"`
	Price       int64      `json:"price"`
	UserID      string     `json:"user_id"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
}

func New(serviceName string, price int64, userID string, startDate time.Time, endDate *time.Time) *Subscription {
	return &Subscription{
		ID:          ID(uuid.NewString()),
		ServiceName: serviceName,
		Price:       price,
		UserID:      userID,
		StartDate:   startDate,
		EndDate:     endDate,
	}
}
