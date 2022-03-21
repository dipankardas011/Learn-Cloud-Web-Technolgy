package test

import (
	"github.com/dipankardas011/banker/helper"
	"testing"
)

func TestDepositAmount(t *testing.T) {

	testDeposit := []float32{-450.24, 0.0, 500.67, 23.0}
	testBalCur := []float32{234.9, 23432, 0}

	for _, ele := range testBalCur {
		for _, deposit := range testDeposit {
			x := helper.AccountHolder{Balance: ele}
			err := x.Deposit(deposit)
			if deposit <= 0 {
				if err == nil {
					t.Fatalf("Deposit is %v is WRONG but the function returned no ERROR\n", deposit)
				}
			} else {
				if err != nil {
					t.Fatalf("Deposit is %v is CORRECT but the function returned ERROR\n", deposit)
				}
				if (deposit + ele) != x.Balance {
					t.Fatalf("(%v + %v) != {%v}\n", deposit, ele, x.Balance)
				}
			}
		}
	}
}

func TestWithdrawAmount(t *testing.T) {

	testWithdraw := []float32{-450.24, 0.0, 500.67, 23.0, -234.23}
	testBalCur := []float32{234.9, 23432, 0, 23.33}

	for _, ele := range testBalCur {
		for _, withDraw := range testWithdraw {
			x := helper.AccountHolder{Balance: ele}
			_, err := x.Withdraw(withDraw)
			if withDraw <= 0 {
				if err == nil {
					t.Fatalf("Withdrawn amount is %v is WRONG but the function returned no ERROR\n", withDraw)
				}
			} else {
				if err != nil {
					if withDraw <= ele {
						t.Fatalf("Withdraw is %v is CORRECT but the function returned ERROR\n", withDraw)
					}
				} else {
					if (ele - withDraw) != x.Balance {
						t.Fatalf("(%v - %v) != {%v}\n", ele, withDraw, x.Balance)
					}
				}
			}
		}
	}
}
