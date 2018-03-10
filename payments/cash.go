package payments

import "fmt"

type Cash struct{}

func CreateCashAccount() *Cash {
	return &Cash{}
}

func (c Cash) ProcessPayment(float32) bool {
	fmt.Println("Processing a cash payment....")
	return true
}
