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

func CreateOrder(directionLong bool, quantity string) (order *futures.CreateOrderResponse) {
	tmpSide := futures.SideTypeBuy
	if directionLong == false {
		tmpSide = futures.SideTypeSell
	}
	order, err := futuresClient.NewCreateOrderService().Symbol(tradePair).
		Side(tmpSide).Type(futures.OrderTypeMarket).Quantity(quantity).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return
}
