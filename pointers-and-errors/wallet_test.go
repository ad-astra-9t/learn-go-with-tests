package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit to a wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("Wallet: %#v, got: %v, want: %v", wallet, got, want)
		}
	})
	t.Run("withdraw from a wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		amount := Bitcoin(5)
		t.Logf("Wallet: %#v, wallet at the start", wallet)

		err := wallet.Withdraw(amount)
		if err != nil {
			t.Fatalf("Wallet: %#v, withdraw amount: %v, cannot withdraw from wallet", wallet, amount)
		}

		got := wallet.Balance()
		want := Bitcoin(5)
		if got != want {
			t.Errorf("Wallet: %#v, got: %v, want: %v", wallet, got, want)
		}
	})
	t.Run("withdraw from a wallet with insuffient balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}
		amount := Bitcoin(10)
		t.Logf("Wallet: %#v, wallet at the start", wallet)

		got := wallet.Withdraw(amount)
		want := ErrInsufficientFunds{&wallet, amount}
		if got == nil {
			t.Fatalf("Wallet: %#v, withdraw amount: %v, no error is returned", wallet, amount)
		}
		if got != want {
			t.Errorf("Wallet: %#v, got: %v, want: %v", wallet, got, want)
		}
	})

}
