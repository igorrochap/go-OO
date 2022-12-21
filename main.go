package main

import (
	"bank/accounts"
	"fmt"
)

func main() {
	// account := Account{holder: "Igor", agencyNumber: 123, accountNumber: 123456, balance: 123.45}
	account := accounts.Account{"Igor", 123, 123456, 300}
	otherAccount := accounts.Account{Holder: "John", Balance: 100}

	status := account.Transfer(250, &otherAccount)
	fmt.Println(status)

	fmt.Println(account)
	fmt.Println(otherAccount)

	// var otherAccount *Account
	// otherAccount = new(Account)
	// otherAccount.holder = "John"
	// otherAccount.balance = 500

	// fmt.Println(*otherAccount)
}
