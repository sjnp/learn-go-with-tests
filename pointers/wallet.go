package wallet

import (
	"errors"
	"fmt"
)

type Btc int

func (btc Btc) String() string {
	return fmt.Sprintf("%d BTC", btc)
}

type Wallet struct {
	balance Btc
}

var ErrInsufficientAmount = errors.New("insufficient amount")

func (w *Wallet) Deposit(amount Btc) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Btc) error {
	if amount > w.balance {
		return ErrInsufficientAmount
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Btc {
	return w.balance
}
