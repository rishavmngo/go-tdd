package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// type Stringer interface {
// 	String() string
// }

func (wallet *Wallet) Deposit(amt Bitcoin) {
	wallet.balance += amt

}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (b Bitcoin) String() string {

	return fmt.Sprintf("%d BTC", b)
}

func (wallet *Wallet) Withdraw(amt Bitcoin) error {
	if amt <= wallet.balance {
		wallet.balance -= amt
		return nil
	}
	return ErrInsufficientFunds
}
