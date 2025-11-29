package query

import (
	"context"

	"shikposh-backend/internal/admin/adapter/repository"
)

type AdminQueryHandler struct {
	adminRepo *repository.AdminRepository
}

func NewAdminQueryHandler(adminRepo *repository.AdminRepository) *AdminQueryHandler {
	return &AdminQueryHandler{
		adminRepo: adminRepo,
	}
}

// GetDashboardStats gets dashboard statistics
func (h *AdminQueryHandler) GetDashboardStats(ctx context.Context) (*repository.DashboardStats, error) {
	return h.adminRepo.GetDashboardStats(ctx)
}

// GetDailySales gets daily sales data
func (h *AdminQueryHandler) GetDailySales(ctx context.Context, days int) ([]repository.DailySalesData, error) {
	return h.adminRepo.GetDailySales(ctx, days)
}
