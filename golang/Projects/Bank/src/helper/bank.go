package helper

import (
	"errors"
	"fmt"
	"time"
)

type BankAccounts struct {
	Customer   AccountHolder
	BankBranch string
}

func (c *AccountHolder) Deposit(deposit float32) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if deposit <= 0 {
		return errors.New("INV deposit")
	}

	fmt.Printf("Deposit amount: %v\t<<<<< %v\n", deposit, time.Now())
	c.Balance += deposit
	return nil
}

func (c *AccountHolder) Withdraw(amount float32) (amtReceived float32, err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	fmt.Printf("Requested amount: %v\t<<<<< %v\n", amount, time.Now())

	if amount > c.Balance || amount <= 0 {
		return 0, errors.New("⚠️Your Current Balance is Insufficient OR withdraw amount is negative")
	}
	amtReceived = amount
	c.Balance -= amount
	return
}

func (fullC *BankAccounts) CheckBalance() {
	fullC.Customer.lock.Lock()
	defer fullC.Customer.lock.Unlock()

	fmt.Printf("@@ YOUR ACCOUNT SUMMARY @@\t<<<<< %v\n", time.Now())
	fmt.Printf("\tNAME: %v\n", fullC.Customer.Name)
	fmt.Printf("\tACCNO.: %v\n", fullC.Customer.AccNumber)
	fmt.Printf("\tAVIL_BAL: %v\n", fullC.Customer.Balance)
	fmt.Printf("\tBRANCH: %v\n", fullC.BankBranch)
	fmt.Printf("@@ ------- @@\n")
}
