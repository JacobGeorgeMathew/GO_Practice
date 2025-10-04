package main

import "fmt"

// Interface
type PaymentProcessor interface {
    Pay(amount float64) string
}

// Implementation 1
type PayPal struct{
    //data int
}

func (p PayPal) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f using PayPal", amount)
}

// Implementation 2
type CreditCard struct{}

func (c CreditCard) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f using Credit Card", amount)
}

type UPI struct{}

func (u UPI) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using UPI",amount)
}

// Function using interface
func Checkout(p PaymentProcessor, amount float64) {
    fmt.Println(p.Pay(amount))
}

func main() {
    Checkout(PayPal{}, 100.50)
    Checkout(CreditCard{}, 200.75)
		Checkout(UPI{},300.50)
}
