package pointers

type Wallet struct {
	currentBalance int
}

func (wallet Wallet) Deposit(amt int) {
	wallet.currentBalance += amt

}

func (wallet Wallet) Balance() int {
	return wallet.currentBalance
}
