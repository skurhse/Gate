package models

type Wallet struct {
	ID                       int
	ReferenceCurrencyBalance float64
	BlockedBalance           float64
	ActiveBalance            float64
	TotalBalance             float64
	Currency                 string
}
