package main

import "fmt"

// Job description
type Payment interface {
    Pay(amount float64)
}

// PayPal
type PayPal struct{}
func (PayPal) Pay(amount float64) {
    fmt.Printf("Paid %.2f using PayPal\n", amount)
}

// CreditCard
type CreditCard struct{}
func (CreditCard) Pay(amount float64) {
    fmt.Printf("Paid %.2f using CreditCard\n", amount)
}

func Checkout(p Payment, amount float64) {
    p.Pay(amount)
}

func main() {
    Checkout(PayPal{}, 99.99)
    Checkout(CreditCard{}, 49.50)
}
