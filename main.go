package main

import "fmt"

type Account struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (account *Account) draw(value float64) string {
	canDraw := value <= account.balance && value > 0
	if canDraw {
		account.balance -= value
		return "Draw successful!"
	}
	return "Draw not allowed!"
}

func main() {
	// account := Account{holder: "Igor", agencyNumber: 123, accountNumber: 123456, balance: 123.45}
	account := Account{"Igor", 123, 123456, 500}

	fmt.Println(account)
	fmt.Println(account.draw(200))
	fmt.Println(account.balance)

	// var otherAccount *Account
	// otherAccount = new(Account)
	// otherAccount.holder = "John"
	// otherAccount.balance = 500

	// fmt.Println(*otherAccount)
}
