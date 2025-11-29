package handler

import (
	"strconv"

	"shikposh-backend/internal/admin/query"
	mw "shikposh-backend/pkg/middleware"

	httpapi "github.com/ali-mahdavi-dev/shikposh-framework/api/http"
	"github.com/ali-mahdavi-dev/shikposh-framework/errors"

	"github.com/gofiber/fiber/v3"
)

type AdminHandler struct {
	adminQueryHandler *query.AdminQueryHandler
	middleware        *mw.Middleware
}

func NewAdminHandler(adminQueryHandler *query.AdminQueryHandler, middleware *mw.Middleware) *AdminHandler {
	return &AdminHandler{
		adminQueryHandler: adminQueryHandler,
		middleware:        middleware,
	}
}

func (h *AdminHandler) RegisterRoutes(r fiber.Router) {
	// Admin routes - require admin or superuser
	// Note: AuthMiddleware should already be registered globally
	adminRoute := r.Group("/api/v1/admin", h.middleware.AdminMiddleware())
	{
		adminRoute.Get("/dashboard/stats", h.GetDashboardStats)
		adminRoute.Get("/dashboard/daily-sales", h.GetDailySales)
	}
}

// GetDashboardStats godoc
//
//	@Summary		Get dashboard statistics
//	@Description	Retrieves overall dashboard statistics (total orders, users, products, revenue)
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	httpapi.ResponseResult	"Dashboard statistics"
//	@Failure		401	{object}	httpapi.ResponseResult	"Unauthorized"
//	@Failure		403	{object}	httpapi.ResponseResult	"Forbidden - Admin access required"
//	@Failure		500	{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/admin/dashboard/stats [get]
func (h *AdminHandler) GetDashboardStats(c fiber.Ctx) error {
	ctx := c.Context()

	stats, err := h.adminQueryHandler.GetDashboardStats(ctx)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, stats)
}

// GetDailySales godoc
//
//	@Summary		Get daily sales data
//	@Description	Retrieves daily sales data for chart visualization (default: last 7 days)
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			days	query		int		false	"Number of days (default: 7)"	example(7)
//	@Success		200		{object}	httpapi.ResponseResult	"Daily sales data"
//	@Failure		401		{object}	httpapi.ResponseResult	"Unauthorized"
//	@Failure		403		{object}	httpapi.ResponseResult	"Forbidden - Admin access required"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/admin/dashboard/daily-sales [get]
func (h *AdminHandler) GetDailySales(c fiber.Ctx) error {
	ctx := c.Context()

	// Get days parameter, default to 7
	daysStr := c.Query("days", "7")
	days, err := strconv.Atoi(daysStr)
	if err != nil || days < 1 || days > 365 {
		return httpapi.ResError(c, errors.Validation("", "days must be between 1 and 365"))
	}

	dailySales, err := h.adminQueryHandler.GetDailySales(ctx, days)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, dailySales)
}
