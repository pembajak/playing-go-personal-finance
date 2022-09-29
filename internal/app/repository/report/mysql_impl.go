package report

import (
	"context"
	"strings"

	"github.com/pembajak/personal-finance/internal/app/models"
	"github.com/pembajak/personal-finance/internal/pkg/driver/driversql"
)

type db struct {
	DB *driversql.Database
}

// NewDB ...
func NewDB(d *driversql.Database) Repository {
	return &db{
		DB: d,
	}
}

// GetTotalTransactionDaily ...
func (d *db) GetTotalTransactionDaily(ctx context.Context, req models.GetTotalTransaction) (returnData []models.TotalTransactionRes, err error) {
	query := d.DB.Instance.WithContext(ctx).
		Table("finances").
		Select(`transaction_date, type, sum(amount) as total`).
		Group("transaction_date").
		Group("type").
		Where("user_id =?", req.UserID)

	if strings.Compare(req.Type, "") != 0 {
		query = query.Where("type =?", req.Type)
	}

	err = query.Find(&returnData).Error
	if err != nil {
		return
	}

	return
}

// GetTotalTransactionMonthly ...
func (d *db) GetTotalTransactionMonthly(ctx context.Context, req models.GetTotalTransaction) (returnData []models.TotalTransactionMonthlyRes, err error) {
	query := d.DB.Instance.WithContext(ctx).
		Table("finances").
		Select(`
			DATE_FORMAT(transaction_date, '%M') as month,
            DATE_FORMAT(transaction_date, '%m') as month_num,
            extract(year from transaction_date) as year,
			type,
            sum(amount) as total
		`).
		Group("month").
		Group("month_num").
		Group("year").
		Group("type").
		Where("user_id =?", req.UserID)

	if strings.Compare(req.Type, "") != 0 {
		query = query.Where("type =?", req.Type)
	}

	err = query.Find(&returnData).Error
	if err != nil {
		return
	}

	return
}
