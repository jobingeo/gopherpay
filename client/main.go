package main

import (
	"../payments"
)

func main() {
	// credit := payments.CreateCreditAccount(
	// 	"Prabhat Bahukhandi",
	// 	"11111111111",
	// 	5,
	// 	2021,
	// 	123)
	// fmt.Println("Address of the instance", &credit, credit)

	// fmt.Println("credit", credit.OwnerName())
	// fmt.Println("credit", credit.CardNumber())

	// fmt.Println("Trying to change card number")
	// err := credit.SetCardNumber("invalid")
	// if err != nil {
	// 	fmt.Println("That didn't work")
	// }
	// fmt.Println("Changed card number:", credit.CardNumber())
	// fmt.Println("Available Credit:", credit.AvailableCredit())

	// Process payments
	var option payments.PaymentOption
	option = payments.CreateCreditAccount(
		"Prabhat Bahukhandi",
		"11111111111",
		5,
		2021,
		123)

	option.ProcessPayment(500)

	option = payments.CreateCashAccount()

	option.ProcessPayment(500)

}
