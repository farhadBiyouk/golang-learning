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

	fullTimeEmployees := []FullTimeEmployee{
		{ID: 1, NationalCode: "1234567890", FullName: "Pejman Rezaee", ExtraHours: 40},
		{ID: 2, NationalCode: "4836524125", FullName: "Maryam Hosseini", ExtraHours: 120},
	}

	partTimeEmployees := []PartTimeEmployee{
		{ID: 3, NationalCode: "6563453455", FullName: "Milad Hassani", PartTimeHours: 100},
		{ID: 4, NationalCode: "5435435435", FullName: "Maryam Rezaee", PartTimeHours: 87},
	}

	shiftEmployees := []ShiftEmployee{
		{ID: 5, NationalCode: "3123123213", FullName: "Shahin", ShiftHours: 150},
		{ID: 6, NationalCode: "6363454355", FullName: "Masoud", ShiftHours: 168},
	}

	// added data in integrate slice
	var employees []Employee = []Employee{}
	for _, item := range fullTimeEmployees {
		employees = append(employees, item)
	}
	for _, item := range partTimeEmployees {
		employees = append(employees, item)
	}
	for _, item := range shiftEmployees {
		employees = append(employees, item)
	}

	// fetch all data from integrated slice

	for _, item := range employees {
		salary, tax := item.SalaryCalculator()
		fmt.Printf("\nEmployee (%T): %v\nSalary: %f\nTax: %f\n", item, item, salary, tax)

	}

}

type Employee interface {
	SalaryCalculator() (salary, tax float64)
}

type FullTimeEmployee struct {
	ID           int
	NationalCode string
	FullName     string
	ExtraHours   float64
}

type PartTimeEmployee struct {
	ID            int
	NationalCode  string
	FullName      string
	PartTimeHours float64
}

type ShiftEmployee struct {
	ID           int
	NationalCode string
	FullName     string
	ShiftHours   float64
}

func (employee FullTimeEmployee) SalaryCalculator() (salary, tax float64) {
	salary = BaseSalary + (ExtraHourRate*employee.ExtraHours)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee PartTimeEmployee) SalaryCalculator() (salary, tax float64) {
	salary = HourlySalaryRate * employee.PartTimeHours
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee ShiftEmployee) SalaryCalculator() (salary, tax float64) {
	salary = ShitSalaryRate * employee.ShiftHours
	tax = salary * TaxRate
	salary += tax
	return
}
