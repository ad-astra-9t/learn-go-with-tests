package main

import (
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

type ErrInsufficientFunds struct {
	w      *Wallet
	amount Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds{w, amount}
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (e ErrInsufficientFunds) Error() string {
	return fmt.Sprintf(
		"Wallet: %#v, withdraw %v with insuffient balance %v",
		*(e.w), e.amount, e.w.balance,
	)
}
