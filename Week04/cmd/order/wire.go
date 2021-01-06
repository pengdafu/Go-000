//+build wireinject

package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"week04/internal/biz"
	"week04/internal/data"
	"week04/internal/service"
)

func InitOrderService(db *gorm.DB) *service.OrderService {
	wire.Build(service.NewOrderService, biz.NewOrderCases, data.NewOrderRepo)
	return &service.OrderService{}
}
