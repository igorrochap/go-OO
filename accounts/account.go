package accounts

type Account struct {
	Holder        string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (account *Account) Draw(value float64) string {
	canDraw := value <= account.Balance && value > 0
	if canDraw {
		account.Balance -= value
		return "Draw successful!"
	}
	return "Draw failed!"
}

func (account *Account) Deposit(value float64) (string, float64) {
	if value > 0 {
		account.Balance += value
		return "Deposit successful!", account.Balance
	}
	return "Deposit failed!", account.Balance
}

func (account *Account) Transfer(value float64, destiny *Account) bool {
	if value < account.Balance && value > 0 {
		account.Balance -= value
		destiny.Deposit(value)
		return true
	}
	return false
}
