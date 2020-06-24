package pointers_and_errors

import (
	"errors"
	"fmt"
)

// A nice way to mask the original type and make it more meaningful.
type Bitcoin int

type Wallet struct{
	balance Bitcoin
}

// I think you can think of this is a ToString() method on a particular type.
func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}

// By default, function args are copied.
// Rather than taking a copy of the w, we take a pointer to the w so we can actually change it.
func (w *Wallet) Balance() Bitcoin{
	return w.balance
}

func (w *Wallet) Deposit(a Bitcoin){
	w.balance += a
}

func (w *Wallet) Withdraw(a Bitcoin) error{

	if !(w.balance > a){
		return errors.New("you've tried to take out more money than you have")
	}

	w.balance -= a
	return nil
}
