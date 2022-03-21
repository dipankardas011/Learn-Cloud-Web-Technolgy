package main

import (
	"fmt"
	"github.com/dipankardas011/banker/helper"
	"math/rand"
)

type DBtype []helper.BankAccounts

var bankAccountData DBtype

func initDatabase() {
	bankAccountData = make(DBtype, 3)
	bankAccountData[0] = helper.BankAccounts{
		Customer: helper.AccountHolder{
			AccNumber: 2000204324,
			Name:      "USR_1",
			Balance:   400.05},
		BankBranch: "PIPEL"}

	bankAccountData[1] = helper.BankAccounts{
		Customer: helper.AccountHolder{
			AccNumber: 2000204326,
			Name:      "USR_4",
			Balance:   40033.30},
		BankBranch: "PIPEL"}

	bankAccountData[2] = helper.BankAccounts{
		Customer: helper.AccountHolder{
			AccNumber: 2000204624,
			Name:      "USR_10",
			Balance:   4234.50},
		BankBranch: "ORIG"}

}

func main() {
	fmt.Println("Welcome To the Bank")

	initDatabase()
	fmt.Printf("Current Number of Users: %v\n", len(bankAccountData))

	for bankTimeWiseCustomer := 0; bankTimeWiseCustomer < 2; bankTimeWiseCustomer++ {
		for customerX := 0; customerX < len(bankAccountData); customerX++ {
			customerr := customerX
			//bankAccountData[customerr].CheckBalance()

			go func() {
				err := bankAccountData[customerr].Customer.Deposit(rand.Float32() * 100)
				if err != nil {
					fmt.Println("Invalid Input")
				}
			}()
			go func() {
				_, err := bankAccountData[customerr].Customer.Withdraw(rand.Float32() * 5000)
				if err != nil {
					fmt.Println("Invalid Input")
				}
			}()

			bankAccountData[customerr].CheckBalance()
		}
	}

}
