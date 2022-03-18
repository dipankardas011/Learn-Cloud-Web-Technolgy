package src

import (
	"fmt"
	"sync"
)

type accountHolder struct {
	name      string
	accNumber int32
	balance   float32
	lock      sync.Mutex
}

type someThing interface{}

func (cust *accountHolder) init(nam string, acc int32, bal float32) {
	fmt.Println("@@ Enter your details @@")
	cust = &accountHolder{
		name: nam,
		accNumber: acc,
		balance: bal
	}
}