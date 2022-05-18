package controller

import (
	"github.com/adshao/go-binance/v2"
)

var (
	apiKey    string
	secretKey string
)

func init() {
	apiKey = GetConfig("API_KEY")
	secretKey = GetConfig("SECRET_KEY")
	client := binance.NewClient(apiKey, secretKey)
	futuresClient := binance.NewFuturesClient(apiKey, secretKey)   // USDT-M Futures
	deliveryClient := binance.NewDeliveryClient(apiKey, secretKey) // Coin-M Futures
}
