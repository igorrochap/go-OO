package accounts

import (
	"bank/clients"
)

type Account struct {
	Holder        clients.Holder
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (account *Account) Draw(value float64) string {
	canDraw := value <= account.balance && value > 0
	if canDraw {
		account.balance -= value
		return "Draw successful!"
	}
	return "Draw failed!"
}

func (account *Account) Deposit(value float64) (string, float64) {
	if value > 0 {
		account.balance += value
		return "Deposit successful!", account.balance
	}
	return "Deposit failed!", account.balance
}

func (account *Account) Transfer(value float64, destiny *Account) bool {
	if value < account.balance && value > 0 {
		account.balance -= value
		destiny.Deposit(value)
		return true
	}
	return false
}

func (account Account) GetBalance() float64 {
	return account.balance
}
