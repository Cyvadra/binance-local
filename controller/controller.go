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

var (
	cacheForm        *FormCreateOrder
	defaultRetString string = "fine"
)

func init() {
	cacheForm = &FormCreateOrder{}
}

func GetOrders(c *gin.Context) {
	currentOrders, err := api.ListOrders()
	if err != nil {
		c.String(200, err.Error())
	} else {
		c.JSON(200, currentOrders)
	}
	return
}

func CloseAllOrders(c *gin.Context) {
	err := api.CloseAllOrders()
	if err != nil {
		c.String(200, err.Error())
	} else {
		c.String(200, defaultRetString)
	}
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
		c.String(200, defaultRetString)
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
	err = api.CreateOrder(isLong, f.Quantity, f.StopLoss, f.TakeProfit)
	if err != nil {
		c.String(200, err.Error())
	}
	c.String(200, defaultRetString)
	return
}

func QueryBalance(c *gin.Context) {
	currentBalance, err := api.QueryBalance()
	if err != nil {
		c.String(200, err.Error())
	} else {
		c.JSON(200, currentBalance)
	}
	return
}
