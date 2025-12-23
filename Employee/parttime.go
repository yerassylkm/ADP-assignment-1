package Employee

type PartTime struct {
	HourlyRate  float64
	HoursWorked int
}

func (pt PartTime) CalculateBonus() float64 {
	if pt.HoursWorked > 40 {
		return float64(pt.HoursWorked-40) * (pt.HourlyRate * 0.5) // Бонус за переработку
	}
	return 0
}

func (pt PartTime) CalculateSalary() float64 {
	return (pt.HourlyRate * float64(pt.HoursWorked)) + pt.CalculateBonus()
}

func (pt PartTime) GetWorkHours() int { return pt.HoursWorked }