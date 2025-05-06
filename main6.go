package main

import (
	"fmt"
)

type makepayment interface {
	deposit(amount float64)
	withdraw(amount float64)
	checkbalance()
}

type bankaccount struct {
	owner   string
	balance float64
}

func (b *bankaccount) deposit(amount float64) {
	if b.balance > 0 {
		fmt.Println("Invalid Amount")
	} else {
		b.balance += amount
		fmt.Println(amount, "deposited successfully. New balance:", b.balance)
	}
}

func (b *bankaccount) withdraw(amount float64) {
	if amount > b.balance {
		fmt.Println("insuffecnt balance")
	} else {
		b.balance -= amount
		fmt.Println(amount, "withdrawed sussesfully.New balance:", b.balance)
	}
}

func (b *bankaccount) checkbalance() {
	fmt.Printf("the balance amount is %v\n", b.balance)
}

func main() {
	var pay makepayment
	pay = &bankaccount{owner: "Amaresha", balance: 1000}

	pay.deposit(0)
	pay.withdraw(200)
	pay.checkbalance()
}
