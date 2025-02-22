package main

import "fmt"

const (
	BaseSalary       = 5_000_000
	ExtraHourRate    = 90_000
	HourlySalaryRate = 110_000
	ShitSalaryRate   = 80_000
	TaxRate          = 0.09
)

func main() {

	employees := []Employee{
		{ID: 1, NationalCode: "1234567890", FullName: "ali rezai", Type: "FullTime", Hour: 40},
		{ID: 2, NationalCode: "1234567891", FullName: "ami hassani", Type: "PartTime", Hour: 120},
		{ID: 3, NationalCode: "1234567892", FullName: "reza alavi", Type: "FullTime", Hour: 15.5},
		{ID: 4, NationalCode: "1234567893", FullName: "gholi razavi", Type: "Shifty", Hour: 15.5},
	}

	for _, item := range employees {
		salary, tax := item.SalaryCalculator(item.Hour)
		fmt.Printf("%s : salary is %f  and tax is  %f \n", item.FullName, salary, tax)
	}
}

type Employee struct {
	ID           int
	NationalCode string
	FullName     string
	Type         string
	Hour         float64
}

func (employee Employee) FullTimeSalaryCalculator(extraHours float64) (salary, tax float64) {
	salary = BaseSalary + (ExtraHourRate*extraHours)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee Employee) PartTimeSalaryCalculator(hours float64) (salary, tax float64) {
	salary = HourlySalaryRate * hours
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee Employee) ShiftSalaryCalculator(hours float64) (salary, tax float64) {
	salary = ShitSalaryRate * hours
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee Employee) SalaryCalculator(hour float64) (salary, tax float64) {
	if employee.Type == "PartTime" {
		return employee.PartTimeSalaryCalculator(hour)
	} else if employee.Type == "FullTime" {
		return employee.FullTimeSalaryCalculator(hour)
	} else {
		return employee.ShiftSalaryCalculator(hour)
	}

}
