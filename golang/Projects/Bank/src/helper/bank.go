package helper

import (
	"fmt"
)

type BankAccounts struct {
	Customer   AccountHolder
	BankBranch string
}

func (c *AccountHolder) Deposit(deposit float32) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.balance += deposit
	return
}

func (c *AccountHolder) Error() string {
	return fmt.Sprintf("Your Current Balance: %v is Insufficient!!!", c.balance)
}

func (c *AccountHolder) Withdraw(amount float32) (amtRecieved float32, err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if amount > c.balance {
		return 0, err
	}
	amtRecieved = amount
	c.balance -= amount
	return
}

func (fullC *BankAccounts) CheckBalance() {
	fullC.Customer.lock.Lock()
	defer fullC.Customer.lock.Unlock()
	fmt.Println("@@ Your Account details @@")
	fmt.Printf("NAME: %v\n", fullC.Customer.name)
	fmt.Printf("ACCNO.: %v\n", fullC.Customer.accNumber)
	fmt.Printf("AVIL_BAL: %v\n", fullC.Customer.balance)
	fmt.Printf("BRANCH: %v\n", fullC.BankBranch)
	fmt.Println("@@ ------- @@")
}
