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
