package data

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"week04/internal/biz"
)

type Order struct {
	ID    int
	Title string
	Price float64
}

func (o Order) TableName() string {
	return "order"
}

type orderRepo struct {
	db *gorm.DB
}

func (o *orderRepo) ListOrders(limit, page int) (res *[]Order, err error) {
	err = o.db.First(res).Limit(limit).Offset(page * limit).Error
	return res, errors.Wrap(err, "ListOrders fail")
}

func NewOrderRepo(db *gorm.DB) biz.OrderRepo {
	return &orderRepo{
		db: db,
	}
}
