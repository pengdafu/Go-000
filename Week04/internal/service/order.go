package service

import (
	"context"
	v1 "week04/api/order/v1"
	"week04/internal/biz"
)

type OrderService struct {
	o *biz.OrderCases
	v1.UnimplementedOrderServiceServer
}

func NewOrderService(o *biz.OrderCases) *OrderService {
	return &OrderService{o: o}
}

func (o *OrderService) ListOrders(ctx context.Context, in *v1.ListOrdersRequest) (*v1.ListOrdersResponse, error) {
	res, err := o.o.ListOrders(&biz.OrderDo{
		Limit: int(in.Limit),
		Page:  int(in.Page),
	})
	if err != nil {
		return nil, err
	}
	r := new(v1.ListOrdersResponse)
	r.Count = uint64(len(*res))
	r.Rows = make([]*v1.OrderResponse, 0)
	for _, order := range *res {
		r.Rows = append(r.Rows, &v1.OrderResponse{
			Id:    uint64(order.Id),
			Name:  order.Name,
			Price: uint64(order.Price),
		})
	}
	return r, nil
}
