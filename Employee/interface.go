package Employee

type Employee interface {
	CalculateSalary() float64
	CalculateBonus() float64
	GetWorkHours() int
}