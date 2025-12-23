package Wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("not enough money in the wallet")

type Wallet struct {
	balance      float64
	transactions []Transaction
}

func (w *Wallet) AddMoney(amount float64) {
	w.balance += amount
	w.transactions = append(w.transactions, Transaction{
		Amount: amount, Type: "Deposit",
	})
}

func (w *Wallet) SpendMoney(amount float64) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	w.transactions = append(w.transactions, Transaction{
		Amount: amount, Type: "Withdrawal",
	})
	return nil
}

func (w *Wallet) GetBalance() float64 {
	return w.balance
}

func (w *Wallet) ShowHistory() {
	fmt.Println("\n--- Transaction History ---")
	for i, t := range w.transactions {
		fmt.Printf("%d. [%s] $%.2f\n", i+1, t.Type, t.Amount)
	}
}