package api

import (
	"context"

	"github.com/adshao/go-binance/v2/futures"
)

type cachePosition struct {
	Direction bool
	Quantity  string
}

var (
	CurrentPosition       *cachePosition
	emptyPositionQuantity string = "0.000"
)

func init() {
	CurrentPosition = &cachePosition{true, emptyPositionQuantity}
}

func ListOrders() (currentOrders []*futures.Order) {
	currentOrders, err := futuresClient.NewListOpenOrdersService().
		Symbol(tradePair).Do(context.Background())
	if err != nil {
		panic(err)
	}
	return
}

func CloseAllOrders() (err error) {
	currentOrders, err := futuresClient.NewListOpenOrdersService().
		Symbol(tradePair).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if len(currentOrders) == 2 && CurrentPosition.Quantity != emptyPositionQuantity {
		err = CreateSimpleOrder(!CurrentPosition.Direction, CurrentPosition.Quantity)
		if err != nil {
			return
		} else {
			CurrentPosition.Quantity = emptyPositionQuantity
		}
	}
	err = futuresClient.NewCancelAllOpenOrdersService().Symbol(tradePair).Do(context.Background())
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
	if err != nil {
		return
	}
	CurrentPosition.Direction = directionLong
	CurrentPosition.Quantity = quantity
	return
}

func CreateSimpleOrder(directionLong bool, quantity string) (err error) {
	tmpSide := futures.SideTypeBuy
	if directionLong == false {
		tmpSide = futures.SideTypeSell
	}
	tmpService := futuresClient.NewCreateOrderService().Symbol(tradePair)
	_, err = tmpService.
		Side(tmpSide).
		Type(futures.OrderTypeMarket).
		Quantity(quantity).
		Do(context.Background())
	return
}
