package api

import (
	binance "github.com/adshao/go-binance/v2"

	"github.com/adshao/go-binance/v2/futures"
)

var (
	apiKey    string
	secretKey string
	tradePair string
)

// client        *binance.Client
var futuresClient *futures.Client

func Init() {
	// client = binance.NewClient(apiKey, secretKey)
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

func SetTradePair(s string) {
	tradePair = s
	return
}
