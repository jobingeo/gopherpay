package payments

import (
	"errors"
	"fmt"
	"time"
)

type PaymentOption interface {
	ProcessPayment(float32) bool
}

type CreditCard struct {
	ownerName       string
	cardNumber      string
	expiryMonth     int
	expiryYear      int
	securityCode    int
	availableCredit float32 // Through webservices || or have the value
}

// CreateCreditAccount Returning a Credit Card Accunt
func CreateCreditAccount(ownerName, cardNumber string, expiryMonth, expiryYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:    ownerName,
		cardNumber:   cardNumber,
		expiryMonth:  expiryMonth,
		expiryYear:   expiryYear,
		securityCode: securityCode,
	}
}

func (c CreditCard) OwnerName() string {
	return c.ownerName
}

// SetOwnerName do SetOwnerName
func (c *CreditCard) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid OwnerName proveded")
	}
	c.ownerName = value
	fmt.Println("Changing name to:", c.ownerName, &c, c)
	return nil
}

func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

func (c *CreditCard) SetCardNumber(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid cardNumber proveded")
	}

	c.cardNumber = value
	return nil
}

func (c CreditCard) ExpirationDate() (int, int) {
	return c.expiryMonth, c.expiryYear
}

func (c *CreditCard) SetExpirationDate(month, year int) error {
	now := time.Now()
	if year < now.Year() || (year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Expiration Date must lie in the future.")
	}
	c.expiryMonth = month
	c.expiryYear = year
	return nil
}

func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

func (c *CreditCard) SetSecurityCode(value int) error {
	if value < 0 {
		return errors.New("Invalid Security Code")
	}
	c.securityCode = value
	return nil
}

func (c CreditCard) AvailableCredit() float32 {
	return 5000.00
}

func (c CreditCard) ProcessPayment(float32) bool {
	fmt.Println("Processing a credit card payment...")
	return true
}
