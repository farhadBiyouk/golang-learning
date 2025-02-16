package main

import "fmt"

func main() {
	var salary float64
	var minSalary float64 = 5_600_000
	var taxPercent float64 = 0
	var knowledgeBase bool = true

	print("enter your salary: ")
	fmt.Scanln(&salary)

	if salary <= minSalary || knowledgeBase {
		taxPercent = 0
	} else if salary > minSalary && salary <= minSalary*2 {
		taxPercent = 0.05
	} else if salary > minSalary && salary <= minSalary*3 {
		taxPercent = 0.07
	} else if salary > minSalary && salary <= minSalary*4 {
		taxPercent = 0.10
	} else {
		taxPercent = 0.15
	}

	fmt.Printf("your tax is: %.2f\n", taxPercent*salary)
	fmt.Printf("your salary is: %.2f\n", salary-taxPercent*salary)

}
