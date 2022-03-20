package helper

import (
	"sync"
)

type AccountHolder struct {
	name      string
	accNumber int32
	balance   float32
	lock      sync.Mutex
}

func (cust *AccountHolder) Init(nam string, acc int32, bal float32) {
	cust.lock.Lock()
	defer cust.lock.Unlock()

	cust.name = nam
	cust.accNumber = acc
	cust.balance = bal
}
