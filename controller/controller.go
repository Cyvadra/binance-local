package controller

import (
	"github.com/Cyvadra/binance-local/api"
	"github.com/gin-gonic/gin"
)

type FormCreateOrder struct {
	Direction  string `form:"direction"`
	Quantity   string `form:"quantity"`
	StopLoss   string `form:"stop_loss"`
	TakeProfit string `form:"take_profit"`
}

var cacheForm *FormCreateOrder

func init() {
	cacheForm = &FormCreateOrder{}
}

func GetOrders(c *gin.Context) {
	c.JSON(200, api.ListOrders())
	return
}

func CloseAllOrders(c *gin.Context) {
	err := api.CloseAllOrders()
	c.String(200, err.Error())
	return
}

func CreateOrder(c *gin.Context) {
	f := &FormCreateOrder{}
	err := c.ShouldBind(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	if cacheForm.Direction == f.Direction && cacheForm.Quantity == f.Quantity {
		c.String(200, "done already")
		return
	} else {
		cacheForm = f
	}
	isLong := true
	if f.Direction == "long" {
		isLong = true
	} else if f.Direction == "short" {
		isLong = false
	} else {
		c.String(500, "direction incorrect")
		return
	}
	c.JSON(200, api.CreateOrder(isLong, f.Quantity, f.StopLoss, f.TakeProfit))
	return
}
