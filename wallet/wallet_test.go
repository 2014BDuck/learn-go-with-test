// @Author: 2014BDuck
// @Date: 2021/2/14

package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Errorf("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(5)
		if err != nil {
			t.Errorf("Got err")
		}
		want := Bitcoin(5)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		originalBalance := Bitcoin(20)
		wallet := Wallet{balance: originalBalance}
		err := wallet.Withdraw(50)
		assertBalance(t, wallet, originalBalance)

		assertError(t, err, InsufficientFundsError)
	})
}
