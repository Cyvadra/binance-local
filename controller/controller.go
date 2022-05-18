package controller

import (
	"github.com/Cyvadra/binance-local/api"
	"github.com/gin-gonic/gin"
)

func GetOpenOrders(c *gin.Context) {
	c.JSON(200, api.ListOpenOrders())
	return
}
