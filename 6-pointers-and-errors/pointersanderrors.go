package pointersanderrors

import (
	"errors"
	"fmt"
)

const INSUFFICIENT_BALANCE_MESSAGE = "insufficient balance"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func New(amount Bitcoin) Wallet {
	return Wallet{
		balance: amount,
	}
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.balance += deposit
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New(INSUFFICIENT_BALANCE_MESSAGE)
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
