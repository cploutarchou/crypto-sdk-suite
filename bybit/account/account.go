package account

import "github.com/cploutarchou/crypto-sdk-suite/bybit/client"

type Account interface {
	Wallet() *Wallet
	UpgradeToUnified() *UpgradeToUnified
	Borrow() *Borrow
	Collateral() *CollateralCoin
	CoinGreek() *CoinGreeks
	FeeRates() *FeeRates
	Info() *Info
	TransactionLog() *TransactionLog
	Margin() *Margin
}

type account struct {
	client *client.Client
}

func (a *account) Collateral() *CollateralCoin {
	return NewSetCollateralCoin(a.client)
}

func (a *account) Wallet() *Wallet {
	return NewWallet(a.client)
}

func (a *account) UpgradeToUnified() *UpgradeToUnified {
	return NewUpgradeToUnifiedRequest(a.client)
}

func (a *account) Borrow() *Borrow {
	return NewBorrow(a.client)
}

func (a *account) CoinGreek() *CoinGreeks {
	return NewCoinGreeks(a.client)
}

func (a *account) FeeRates() *FeeRates {
	return NewFeeRates(a.client)
}

func (a *account) Info() *Info {
	return NewInfo(a.client)
}
func (a *account) TransactionLog() *TransactionLog {
	return NewTransactionLog(a.client)
}
func (a *account) Margin() *Margin {
	return NewMargin(a.client)
}
func New(client *client.Client) Account {
	return &account{client: client}
}
