package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

// DashboardStats represents dashboard statistics
type DashboardStats struct {
	TotalOrders   int64 `json:"total_orders"`
	TotalUsers    int64 `json:"total_users"`
	TotalProducts int64 `json:"total_products"`
	TotalRevenue  int64 `json:"total_revenue"`
}

// DailySalesData represents daily sales for chart
type DailySalesData struct {
	Date   string `json:"date"`
	Sales  int64  `json:"sales"`
	Orders int    `json:"orders"`
}

// GetDashboardStats gets overall dashboard statistics
func (r *AdminRepository) GetDashboardStats(ctx context.Context) (*DashboardStats, error) {
	stats := &DashboardStats{}

	// Get total orders
	if err := r.db.WithContext(ctx).
		Table("orders").
		Where("deleted_at IS NULL").
		Count(&stats.TotalOrders).Error; err != nil {
		return nil, err
	}

	// Get total users
	if err := r.db.WithContext(ctx).
		Table("users").
		Where("deleted_at IS NULL").
		Count(&stats.TotalUsers).Error; err != nil {
		return nil, err
	}

	// Get total products
	if err := r.db.WithContext(ctx).
		Table("products").
		Where("deleted_at IS NULL").
		Count(&stats.TotalProducts).Error; err != nil {
		return nil, err
	}

	// Get total revenue (sum of final_amount from orders with payment_confirmed or delivered status)
	var revenue int64
	if err := r.db.WithContext(ctx).
		Table("orders").
		Where("deleted_at IS NULL AND status IN (?, ?, ?)", "payment_confirmed", "confirmed", "delivered").
		Select("COALESCE(SUM(final_amount), 0)").
		Scan(&revenue).Error; err != nil {
		return nil, err
	}
	stats.TotalRevenue = revenue

	return stats, nil
}

// GetDailySales gets daily sales data for the last N days
func (r *AdminRepository) GetDailySales(ctx context.Context, days int) ([]DailySalesData, error) {
	var results []struct {
		Date   time.Time `gorm:"column:date"`
		Sales  int64     `gorm:"column:sales"`
		Orders int       `gorm:"column:orders"`
	}

	// Calculate start date
	startDate := time.Now().AddDate(0, 0, -days)

	// Query to get daily sales grouped by date
	query := `
		SELECT 
			DATE(created_at) as date,
			COALESCE(SUM(final_amount), 0) as sales,
			COUNT(*) as orders
		FROM orders
		WHERE deleted_at IS NULL 
			AND created_at >= ?
			AND status IN ('payment_confirmed', 'confirmed', 'delivered')
		GROUP BY DATE(created_at)
		ORDER BY date ASC
	`

	if err := r.db.WithContext(ctx).Raw(query, startDate).Scan(&results).Error; err != nil {
		return nil, err
	}

	// Convert to DailySalesData format
	dailySales := make([]DailySalesData, 0, len(results))
	for _, result := range results {
		// Format date as Persian date string
		dateStr := result.Date.Format("2006-01-02")
		dailySales = append(dailySales, DailySalesData{
			Date:   dateStr,
			Sales:  result.Sales,
			Orders: result.Orders,
		})
	}

	// Fill in missing days with zero values
	completeData := r.fillMissingDays(dailySales, days)

	return completeData, nil
}

// fillMissingDays fills missing days with zero values
func (r *AdminRepository) fillMissingDays(data []DailySalesData, days int) []DailySalesData {
	// Create a map of existing dates
	dateMap := make(map[string]DailySalesData)
	for _, d := range data {
		dateMap[d.Date] = d
	}

	// Generate all dates for the last N days
	completeData := make([]DailySalesData, 0, days)
	today := time.Now()

	for i := days - 1; i >= 0; i-- {
		date := today.AddDate(0, 0, -i)
		dateStr := date.Format("2006-01-02")

		if existing, ok := dateMap[dateStr]; ok {
			completeData = append(completeData, existing)
		} else {
			completeData = append(completeData, DailySalesData{
				Date:   dateStr,
				Sales:  0,
				Orders: 0,
			})
		}
	}

	return completeData
}
