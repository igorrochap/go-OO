package main

import (
	"bank/accounts"
	"bank/clients"
	"fmt"
)

func main() {
	holder1 := clients.Holder{Name: "Igor", Cpf: "123.123.123-12", Profession: "Software Engineer"}
	account1 := accounts.Account{Holder: holder1, AgencyNumber: 123, AccountNumber: 123456}

	account1.Deposit(150)

	fmt.Println(account1.GetBalance())
}
