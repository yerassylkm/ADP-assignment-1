package Employee

type Contractor struct {
	RatePerProject    float64
	ProjectsCompleted int
}

func (c Contractor) CalculateBonus() float64 {
	return float64(c.ProjectsCompleted) * 150.0 // Фиксированная премия за проект
}

func (c Contractor) CalculateSalary() float64 {
	return (c.RatePerProject * float64(c.ProjectsCompleted)) + c.CalculateBonus()
}

func (c Contractor) GetWorkHours() int { return c.ProjectsCompleted * 20 }