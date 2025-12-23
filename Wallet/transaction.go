package Wallet

import "time"

type Transaction struct {
	Amount    float64
	Timestamp time.Time
	Type      string // "Deposit" or "Withdrawal"
}