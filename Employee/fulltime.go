package Employee

type FullTime struct {
	BaseSalary float64
	BonusPct   float64
}

func (ft FullTime) CalculateBonus() float64  { return ft.BaseSalary * ft.BonusPct }
func (ft FullTime) CalculateSalary() float64 { return ft.BaseSalary + ft.CalculateBonus() }
func (ft FullTime) GetWorkHours() int        { return 160 }