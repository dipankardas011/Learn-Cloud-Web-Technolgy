package main

import "fmt"

type Bank struct {
	name    string
	balance int32
	accNo   uint64
}

func withdraw(acc Bank, amt int32) bool {
	if acc.balance < amt {
		return false
	}
	return true
}

func main() {
	var acc Bank
	fmt.Println("Enter Your name, balance, accNo")
	acc.name = "Dipankar"
	acc.balance = 2342
	acc.accNo = 2342343234
	fmt.Printf("Name: %v\nAccNo: %v\nBalance: %v\n", acc.name, acc.accNo, acc.balance)
	fmt.Println("Can withdraw 1000? ", withdraw(acc, 1000))
	fmt.Println("Can withdraw 10000? ", withdraw(acc, 10000))
}
