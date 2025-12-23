package Gym

import "fmt"

type Member interface {
	GetDetails() string
}

type BasicMember struct {
	ID   uint64
	Name string
}

func (b BasicMember) GetDetails() string {
	return fmt.Sprintf("[ID: %d] Basic Member: %s (Access: Gym Floor Only)", b.ID, b.Name)
}

type PremiumMember struct {
	ID   uint64
	Name string
}

func (p PremiumMember) GetDetails() string {
	return fmt.Sprintf("[ID: %d] Premium Member: %s (Access: Gym, Pool, Sauna)", p.ID, p.Name)
}