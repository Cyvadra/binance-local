package api

import (
	"context"

	binance "github.com/adshao/go-binance/v2"

	"github.com/adshao/go-binance/v2/futures"
)

var (
	apiKey    string
	secretKey string
	tradePair string = "BTCUSDT"
)

var (
	client        *binance.Client
	futuresClient *futures.Client
)

func Init() {
	client = binance.NewClient(apiKey, secretKey)
	futuresClient = binance.NewFuturesClient(apiKey, secretKey) // USDT-M Futures
}

func SetAPIKey(s string) {
	apiKey = s
	return
}

func SetSecretKey(s string) {
	secretKey = s
	return
}