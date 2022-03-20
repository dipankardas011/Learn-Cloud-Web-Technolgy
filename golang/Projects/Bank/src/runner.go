package main

import (
	"fmt"
	"github.com/dipankardas011/banker/helper"
)

func initDatabase() {

}

func main() {
	fmt.Println("Welcome To the Bank")
	fmt.Println("Enter the name, acc number & balance")
	na := ""
	var acc int32
	var bal float32
	fmt.Scanln(&na)

	fmt.Scanln(&acc)
	fmt.Scanln(&bal)

	// TODO nil pointer
	var x *helper.AccountHolder

	x.Init(na, acc, bal)
	//var bankDataBase helper.BankAccounts
	bankDataBase := helper.BankAccounts{Customer: *x, BankBranch: "sd"}

	bankDataBase.CheckBalance()
}
