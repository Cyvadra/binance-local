package api

import (
	"context"

	"github.com/adshao/go-binance/v2/futures"
)

func ListOrders() (openOrders []*futures.Order) {
	openOrders, err := futuresClient.NewListOrdersService().Symbol(tradePair).Do(context.Background())
	if err != nil {
		panic(err)
	}
	return
}

func CreateOrder(directionLong bool, quantity, stopLoss, takeProfit string) (err error) {
	tmpSide := futures.SideTypeBuy
	anoSide := futures.SideTypeSell
	if directionLong == false {
		tmpSide = futures.SideTypeSell
		anoSide = futures.SideTypeBuy
	}
	tmpService := futuresClient.NewCreateOrderService().Symbol(tradePair)
	_, err = tmpService.
		Side(tmpSide).
		Type(futures.OrderTypeMarket).
		Quantity(quantity).
		Do(context.Background())
	if err != nil {
		return
	}
	// stopLoss
	_, err = tmpService.
		Side(anoSide).
		Type(futures.OrderTypeStopMarket).
		WorkingType(futures.WorkingTypeMarkPrice).
		PositionSide(futures.PositionSideTypeBoth).
		ClosePosition(true).
		StopPrice(stopLoss).
		Do(context.Background())
	if err != nil {
		return
	}
	// takeProfit
	_, err = tmpService.
		Side(anoSide).
		Type(futures.OrderTypeTakeProfitMarket).
		WorkingType(futures.WorkingTypeMarkPrice).
		PositionSide(futures.PositionSideTypeBoth).
		ClosePosition(true).
		StopPrice(takeProfit).
		Do(context.Background())
	return
}
