package main

import (
	"fmt"
	"strconv"
)

// --------------

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough cash")
	}
	w.Cash -= amount
	return nil
}

func (w *Wallet) String() string {
	return "Wallet with " + strconv.Itoa(w.Cash) + " money"
}

// --------------

type Payer interface {
	Pay(int) error
}

// --------------

func Buy(in interface{}) {
	var p Payer
	var ok bool
	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T isn't a means of payment\n\n", in)
		return
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Payment error %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Thank you for purchasing via %T\n\n", p)

}

// --------------

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)
	Buy([]int{1, 2, 3})
	Buy(3.14)
}
