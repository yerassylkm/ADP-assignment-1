package main

import (
	"fmt"
	"os"

	"github.com/YerassylKuanyshov/adp-asg1/Employee"
	"github.com/YerassylKuanyshov/adp-asg1/Gym"
	"github.com/YerassylKuanyshov/adp-asg1/Hotel"
	"github.com/YerassylKuanyshov/adp-asg1/Wallet"
)

func main() {
	// Инициализация систем
	hotelSystem := Hotel.NewHotel()
	gymSystem := Gym.NewGym()
	userWallet := &Wallet.Wallet{}

	for {
		fmt.Println("\n========================================")
		fmt.Println("   ASssignnment 1")
		fmt.Println("========================================")
		fmt.Println("1. Hotel Management (Task 1)")
		fmt.Println("2. Employee Payroll Demo (Task 2)")
		fmt.Println("3. Gym Membership (Task 3)")
		fmt.Println("4. Digital Wallet (Task 4)")
		fmt.Println("5. Exit")
		fmt.Print("Choose task to run: ")

		var mainChoice int
		fmt.Scan(&mainChoice)

		switch mainChoice {
		case 1:
			runHotelMenu(hotelSystem)
		case 2:
			runEmployeeDemo()
		case 3:
			runGymDemo(gymSystem)
		case 4:
			runWalletMenu(userWallet)
		case 5:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option.")
		}
	}
}

// --- Task 1: Hotel Logic ---
func runHotelMenu(h *Hotel.HotelManager) {
	fmt.Println("\n--- Hotel Menu ---")
	fmt.Println("1. Add Room\n2. Check-In\n3. List Vacant")
	var c int
	fmt.Scan(&c)
	switch c {
	case 1:
		var id, cat string
		var price float64
		fmt.Print("Room ID: "); fmt.Scan(&id)
		fmt.Print("Category: "); fmt.Scan(&cat)
		fmt.Print("Price: "); fmt.Scan(&price)
		h.AddRoom(Hotel.Room{RoomNumber: id, Type: cat, PricePerNight: price})
	case 2:
		var id string
		fmt.Print("Room ID to check-in: "); fmt.Scan(&id)
		if err := h.CheckIn(id); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println("Success!")
		}
	case 3:
		rooms := h.GetVacantRooms()
		for _, r := range rooms {
			fmt.Printf("Room %s (%s) - $%.2f\n", r.RoomNumber, r.Type, r.PricePerNight)
		}
	}
}

// --- Task 2: Employee Logic (Interface Demo) ---
func runEmployeeDemo() {
	fmt.Println("\n--- Employee Payroll (Polymorphism Demo) ---")
	// Создаем слайс интерфейсов Employee
	staff := []Employee.Employee{
		Employee.FullTime{BaseSalary: 5000, BonusPct: 0.1},
		Employee.PartTime{HourlyRate: 30, HoursWorked: 50},
		Employee.Contractor{RatePerProject: 1500, ProjectsCompleted: 2},
		Employee.Intern{Stipend: 40, Days: 15},
	}

	for _, e := range staff {
		fmt.Printf("Type: %T | Salary: $%.2f | Bonus: $%.2f | Hours: %d\n", 
			e, e.CalculateSalary(), e.CalculateBonus(), e.GetWorkHours())
	}
}

// --- Task 3: Gym Logic ---
func runGymDemo(g *Gym.Gym) {
	fmt.Println("\n--- Gym Management ---")
	g.AddMember(1, Gym.BasicMember{ID: 1, Name: "John Doe"})
	g.AddMember(2, Gym.PremiumMember{ID: 2, Name: "Alice Smith"})
	g.ListMembers()
}

// --- Task 4: Wallet Logic ---
func runWalletMenu(w *Wallet.Wallet) {
	for {
		fmt.Printf("\nCurrent Balance: $%.2f\n", w.GetBalance())
		fmt.Println("1. Deposit\n2. Spend\n3. History\n4. Back")
		var c int
		fmt.Scan(&c)
		if c == 4 { break }
		
		switch c {
		case 1:
			var amt float64
			fmt.Print("Amount: "); fmt.Scan(&amt)
			w.AddMoney(amt)
		case 2:
			var amt float64
			fmt.Print("Amount: "); fmt.Scan(&amt)
			if err := w.SpendMoney(amt); err != nil {
				fmt.Println("Error:", err)
			}
		case 3:
			w.ShowHistory()
		}
	}
}