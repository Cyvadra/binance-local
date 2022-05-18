package api

import (
	"context"

	binance "github.com/adshao/go-binance/v2"

	"github.com/adshao/go-binance/v2/futures"
)

func ListOpenOrders() (openOrders []*binance.Order) {
	openOrders, err := client.NewListOpenOrdersService().Symbol(tradePair).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return
}
