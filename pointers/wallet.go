package pointers
import "fmt"

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// 逆参照不要
	w.balance += amount
}

func (w *Wallet) Ballance() Bitcoin {
	// 逆参照不要
	return w.balance
}

