package main

import "fmt"



func main() {
    
    Checkout(PayPal{}, 99.99)
    Checkout(CreditCard{}, 49.50)
}

type Payment interface {
    Pay(amount float64)
}


type PayPal struct{}
func (PayPal) Pay(amount float64) {

    fmt.Printf("Paid %f using PayPal\n", amount)
}

type CreditCard struct{}

func (CreditCard) Pay(amount float64) {
    
    fmt.Printf("Paid %f using CreditCard\n", amount)
}

func Checkout(p Payment, amount float64) {

    p.Pay(amount)
}
