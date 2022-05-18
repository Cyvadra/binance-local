package api

import (
	"context"

	binance "github.com/adshao/go-binance/v2"
	// "github.com/adshao/go-binance/v2/futures"
)

func ListOpenOrders() (openOrders []*binance.Order) {
	openOrders, err := client.NewListOpenOrdersService().Symbol(tradePair).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return
}

func CreateOrder(directionLong bool, quantity string) (order *binance.CreateOrderResponse) {
	tmpSide := binance.SideTypeBuy
	if directionLong == false {
		tmpSide = binance.SideTypeSell
	}
	order, err := client.NewCreateOrderService().Symbol(tradePair).
		Side(tmpSide).Type(binance.OrderTypeMarket).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity(quantity).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return
}
