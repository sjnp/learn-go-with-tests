package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	walletTests := []struct {
		name                  string
		wallet                Wallet
		depositAmount         Bitcoin
		withdrawAmount        Bitcoin
		expectedWithdrawError error
		expectedBalance       Bitcoin
	}{
		{name: "deposit", wallet: Wallet{}, depositAmount: 10, expectedBalance: 10},
		{name: "withdraw", wallet: Wallet{}, depositAmount: 10, withdrawAmount: 5, expectedBalance: 5},
		{name: "invalid withdraw", wallet: Wallet{}, depositAmount: 10, withdrawAmount: 15, expectedWithdrawError: ErrInsufficientFunds},
	}

	for _, wt := range walletTests {
		t.Run(wt.name, func(t *testing.T) {
			wt.wallet.Deposit(wt.depositAmount)
			err := wt.wallet.Withdraw(wt.withdrawAmount)
			if err != nil || wt.expectedWithdrawError != nil {
				if err == nil {
					t.Errorf("got nil error but expected %q", wt.expectedWithdrawError)
					return
				}

				if wt.expectedWithdrawError == nil {
					t.Errorf("got %q but expected none", err)
					return
				}

				if err.Error() != wt.expectedWithdrawError.Error() {
					t.Errorf("got %q but expected %q", err, wt.expectedWithdrawError)
					return
				}

				return
			}

			result := wt.wallet.Balance()

			if result != wt.expectedBalance {
				t.Errorf("got %s but expected %s from %#v", result, wt.expectedBalance, wt.wallet)
			}
		})
	}
}
