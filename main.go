package main

import (
	"bank/accounts"
	"fmt"
)

func main() {
	savingsAccount := accounts.SavingsAccount{}
	savingsAccount.Deposit(100)

	payTicket(&savingsAccount, 60)

	fmt.Println(savingsAccount.GetBalance())
}

func payTicket(account accounts.Account, value float64) {
	account.Draw(value)
}
