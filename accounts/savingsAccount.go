package accounts

import "bank/clients"

type SavingsAccount struct {
	Holder                                 clients.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (account *SavingsAccount) Draw(value float64) string {
	canDraw := value <= account.balance && value > 0
	if canDraw {
		account.balance -= value
		return "Draw successful!"
	}
	return "Draw failed!"
}

func (account *SavingsAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		account.balance += value
		return "Deposit successful!", account.balance
	}
	return "Deposit failed!", account.balance
}

func (account SavingsAccount) GetBalance() float64 {
	return account.balance
}
