package main

import (
	"fmt"
	"strings"

	"github.com/Cyvadra/binance-local/controller"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	g := gin.Default()
	// functions
	g.GET("/getOrders", controller.GetOrders)
	g.GET("/closeAllOrders", controller.CloseAllOrders)
	g.GET("/createOrder", controller.CreateOrder)
	// empty returns
	g.GET("/", func(c *gin.Context) { c.String(200, "") })
	g.GET("/robots.txt", func(c *gin.Context) { c.String(200, "User-agent: *\nDisallow: *") })
	g.GET("/sitemap.xml", func(c *gin.Context) { c.String(200, "") })
	g.GET("/favicon.ico", func(c *gin.Context) { c.String(404, "") })
	default404str := strings.Repeat("00000000", 1024)
	g.NoRoute(gzip.Gzip(gzip.DefaultCompression), func(c *gin.Context) { c.String(200, default404str) })
	// pull up service
	g.Run("127.0.0.1:8023")
	fmt.Println("server up")
}
