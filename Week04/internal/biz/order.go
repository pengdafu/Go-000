package biz

import (
	"github.com/pkg/errors"
	"week04/internal/data"
)

type OrderRepo interface {
	ListOrders(limit, page int) (*[]data.Order, error)
}

type OrderCases struct {
	o OrderRepo
}

func NewOrderCases(o OrderRepo) *OrderCases {
	return &OrderCases{
		o: o,
	}
}

type OrderDo struct {
	Limit int
	Page  int
	Id    int
	Name  string
	Price float64
}

func (oc *OrderCases) ListOrders(do *OrderDo) (*[]OrderDo, error) {
	orders, err := oc.o.ListOrders(do.Limit, do.Page)
	if err != nil {
		return nil, errors.Wrap(err, "could not list order")
	}
	res := make([]OrderDo, 0)
	for _, order := range *orders {
		res = append(res, OrderDo{
			Id:    order.ID,
			Name:  order.Title,
			Price: order.Price,
		})
	}
	return &res, nil
}
