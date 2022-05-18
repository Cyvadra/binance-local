package controller

import (
	"github.com/Cyvadra/binance-local/api"
	"github.com/gin-gonic/gin"
)

func GetOpenOrders(c *gin.Context) {
	c.JSON(200, api.ListOpenOrders())
	return
}

type FormCreateOrder struct {
	Direction string `form:"direction"`
	Quantity  string `form:"quantity"`
}

func CreateOrder(c *gin.Context) {
	f := &FormCreateOrder{}
	c.ShouldBind(f)
	isLong := true
	if f.Direction == "long" {
		isLong = true
	} else if f.Direction == "short" {
		isLong = false
	} else {
		panic("direction incorrect")
	}
	c.JSON(200, api.CreateOrder(isLong, f.Quantity))
	return
}
