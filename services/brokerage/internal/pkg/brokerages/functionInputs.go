package brokerages

import (
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
	"time"
)

type LoginParams struct {
	Totp int
}

type OrderBookParams struct {
	Symbol repository.Market
}

type MarketStatusParams struct {
	Destination string
	Source      string
}

type OHLCParams struct {
	Resolution *repository.Resolution
	Market     *repository.Market
	From       time.Time
	To         time.Time
}

type WalletInfoParams struct {
	WalletName string
}

type WalletBalanceParams struct {
	Currency string
}

type TransactionListParams struct {
	WalletID int
}

//type NewOrderParams struct {
//	OrderKind  repository.OrderKind
//	ClientUUID string
//	BuyOrSell  repository.OrderType
//	Price      float64
//	StopPrice  float64
//	Market     repository.Market
//	Amount     float64
//	Option     repository.OrderOption
//	HideOrder  bool
//}

type CancelOrderParams struct {
	ServerOrderId int64
	Market        repository.Market
	IsBuy         bool
	ClientUUID    string
	AllOrders     bool
}

type OrderStatusParams struct {
	ServerOrderId int64
	Market        repository.Market
	ClientUUID    string
}

//type OrderListParams struct {
//	WithDetails bool
//	Status      repository.OrderStatus
//	Market      repository.Market
//	Type        repository.OrderType
//	IsBuy       repository.OrderType
//	ClientUUID  uuid.UUID
//	Page        int
//	Limit       int
//	IsExecuted  bool
//}
//
//type UpdateOrderStatusParams struct {
//	NewStatus repository.OrderStatus
//	OrderId   uint64
//}
//
//type MarketInfoParams struct {
//	MarketName string
//}