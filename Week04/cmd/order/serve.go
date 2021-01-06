package main

import (
	"gorm.io/gorm"
	"log"
)

func main() {
	orderService := InitOrderService(&gorm.DB{})
	log.Println(orderService)
}
