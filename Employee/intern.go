package Employee

type Intern struct {
	Stipend float64
	Days    int
}

func (i Intern) CalculateBonus() float64  { return 0 }
func (i Intern) CalculateSalary() float64 { return i.Stipend * float64(i.Days) }
func (i Intern) GetWorkHours() int        { return i.Days * 8 }