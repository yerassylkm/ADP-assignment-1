package Gym

import "fmt"

type Gym struct {
	Members map[uint64]Member
}

func NewGym() *Gym {
	return &Gym{Members: make(map[uint64]Member)}
}

func (g *Gym) AddMember(id uint64, m Member) {
	g.Members[id] = m
}

func (g *Gym) ListMembers() {
	if len(g.Members) == 0 {
		fmt.Println("No members registered yet.")
		return
	}
	fmt.Println("\n--- Current Gym Members ---")
	for _, m := range g.Members {
		fmt.Println(m.GetDetails())
	}
}