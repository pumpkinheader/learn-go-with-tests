package pointers
import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	// 逆参照不要
	w.balance += amount
}

func (w *Wallet) Ballance() int {
	// 逆参照不要
	return w.balance
}

