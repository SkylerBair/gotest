package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}
type Bitcoin int

type Stringer interface {
	String() string
}

assertError := func(t testing.TB, got error, want string) {
	t.Helper()
	if err == nil {
		t.Fatal("Didnt get a error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

t.Run("Withdraw insufficent funds", func(t *testing.T) {
	startingBalance := Bitcoin(20)
	wallet := Wallet{startingBalance}
	err := wallet.Withdraw(Bitcoin(100))

	assertError(t, err, "cannot withdraw, insufficient funds")
	assertBalance(t, wallet, startingBalance)
})

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {

	w.balance += amount

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("oh no")
	}
	w.balance -= amount
	return nil
}


