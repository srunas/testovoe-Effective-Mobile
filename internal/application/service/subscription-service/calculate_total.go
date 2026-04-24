package subscription_service

import (
	"context"
	"log/slog"

	"github.com/srunas/effective-mobile/internal/domain/service"
)

func (s *Implementation) CalculateTotal(
	ctx context.Context,
	req service.CalculateTotalRequest,
) (service.CalculateTotalResponse, error) {
	total, err := s.repo.CalculateTotal(ctx, req.UserID, req.ServiceName, req.From, req.To)
	if err != nil {
		slog.ErrorContext(ctx, "ошибка подсчёта суммы подписок", "error", err)
		return service.CalculateTotalResponse{}, err
	}

	slog.InfoContext(ctx, "сумма подписок подсчитана", "total", total)

	return service.CalculateTotalResponse{Total: total}, nil
}
