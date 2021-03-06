package models

import (
	"time"
)

type Order struct {
	ClientUUID       string `gorm:"primarykey"`
	ServerOrderId    int64
	Amount           float64
	CreatedAt        time.Time
	FinishedAt       time.Time
	ExecutedAmount   float64
	UnExecutedAmount float64
	Status           OrderStatus
	ExecutedPrice    float64
	Market           Market
	MakerFeeRate     float64
	TakerFeeRate     float64
	SellOrBuy        OrderType
	OrderKind        OrderKind
	AveragePrice     float64
	TransactionFee   float64
	SourceAsset      Asset
	DestinationAsset Asset
	FeeAsset         Asset
	FeeDiscount      float64
	AssetFee         float64
	MoneyFee         float64
	StockFee         float64
	User             string

	//Fee                 float64
	//Qty                 int
	//User                string
	//Price               float64
	//TotalPrice          string
	//MatchedVolume       string
	//UnMatchedVolume     string
}

type OrderType string
type OrderKind string
type OrderOption string
type OrderStatus string
type Asset string

const (
	Ask  OrderType = "ask"
	Bid  OrderType = "bid"
	Buy  OrderType = "buy"
	Sell OrderType = "sell"
)

const (
	LimitOrderKind         OrderKind = "limit"
	MarketOrderKind        OrderKind = "market"
	MultipleLimitOrderKind OrderKind = "batch"
	IOCOrderKind           OrderKind = "ioc"
	StopLimitOrderKind     OrderKind = "stop-limit"
	AllOrderKind           OrderKind = "all"
)

const (
	New            OrderStatus = "new"
	Done           OrderStatus = "done"
	PartlyExecuted OrderStatus = "part_deal"
	UnExecuted     OrderStatus = "not_deal"
)

const (
	OptionIOC       OrderOption = "IOC"
	OptionFOK       OrderOption = "FOK"
	OptionNormal    OrderOption = "NORMAL"
	OptionMakerOnly OrderOption = "MAKER_ONLY"
)
