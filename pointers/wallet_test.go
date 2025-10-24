package wallet

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(tb testing.TB, wallet Wallet, expected Btc) {
		tb.Helper()

		got := wallet.Balance()

		if got != expected {
			t.Errorf("got %s but expected %s", got, expected)
		}
	}

	assertError := func(tb testing.TB, got, expected error) {
		t.Helper()
		if got == nil {
			t.Fatal("no error but expected one")
		}

		if !errors.Is(got, expected) {
			t.Errorf("got %q but expected %q", got, expected)
		}
	}

	assertNotError := func(tb testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("have an error but not expected one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Btc(10))

		expected := Btc(10)
		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Btc(20)}

		err := wallet.Withdraw(10)

		expected := Btc(10)

		assertNotError(t, err)
		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{}

		err := wallet.Withdraw(10)

		assertError(t, err, ErrInsufficientAmount)
	})
}
