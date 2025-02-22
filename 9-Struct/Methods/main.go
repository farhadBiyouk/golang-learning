package main

import "fmt"

type Employee struct {
	ID            int
	Name          string
	Type          string
	SalaryOfMonth int
}

func main() {

	emp1 := &Employee{ID: 1, Name: "Ali", Type: "Manger", SalaryOfMonth: 1000}
	emp2 := &Employee{ID: 1, Name: "Amir", Type: "Programmer", SalaryOfMonth: 900}
	// call function
	fmt.Println(CalculationSalaryInYear(*emp1))
	// call method for struct employee
	fmt.Println(emp1.CalculationSalaryInYear())
	emp1.GetInfo()
	fmt.Println(emp2.CalculationSalaryInYear())
	emp2.GetInfo()

}

// function
func CalculationSalaryInYear(employee Employee) int {

	return employee.SalaryOfMonth * 12
}

// method
func (employee *Employee) CalculationSalaryInYear() int {

	return employee.SalaryOfMonth * 12
}

// method
func (employee Employee) GetInfo() {
	fmt.Printf("ID: %d {%s - %s - %d} \n", employee.ID, employee.Name, employee.Type, employee.SalaryOfMonth)
}
