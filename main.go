package main

import "fmt"

type Account struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	// account := Account{holder: "Igor", agencyNumber: 123, accountNumber: 123456, balance: 123.45}
	account := Account{"Igor", 123, 123456, 123.45}
	fmt.Println(account)
}
