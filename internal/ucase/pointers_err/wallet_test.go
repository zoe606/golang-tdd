package pointers_err

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		expected := Bitcoin(10)
		assertBalance(t, wallet, expected)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}
		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			return
		}
		expected := Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, wallet, expected)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, expected error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != expected {
		t.Errorf("got %s expected %s", got, expected)
	}
}
