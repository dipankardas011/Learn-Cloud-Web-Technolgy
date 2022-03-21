package helper

import (
	"sync"
)

type AccountHolder struct {
	AccNumber int32
	Name      string
	Balance   float32
	lock      sync.Mutex
}

func (cust *AccountHolder) Init(nam string, acc int32, bal float32) {
	cust.lock.Lock()
	defer cust.lock.Unlock()

	cust.Name = nam
	cust.AccNumber = acc
	cust.Balance = bal
}
