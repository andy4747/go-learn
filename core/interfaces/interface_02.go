package interfaces

import "fmt"

//SalaryCalculator interface
type SalaryCalculator interface {
	CalculateSalary() int
}

//PermanentEmployee struct stores the salary data of permemployees
type PermanentEmployee struct {
	ID     int
	salary int
	pf     int
}

//TemporaryEmployee struct stores the salary data of temp employees
type TemporaryEmployee struct {
	ID     int
	salary int
}

//Freelancer struct stores the salary data of freelancers employees
type Freelancer struct {
	ID          int
	ratePerHour int
	totalHours  int
}

//CalculateSalary method to calculate the total salary of perm employee
func (pe PermanentEmployee) CalculateSalary() int {
	return pe.salary + pe.pf
}

//CalculateSalary method to calculate the total salary of temp employee
func (te TemporaryEmployee) CalculateSalary() int {
	return te.salary
}

//CalculateSalary method to calculate the total salary of freelancers employee
func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func totalExpenses(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}
	fmt.Println("Total Expense Per Month: ", expense)
}

//Interface2 is second interface method during practice
func Interface2() {
	p1 := PermanentEmployee{
		ID:     1,
		salary: 1000,
		pf:     10,
	}
	p2 := PermanentEmployee{
		ID:     2,
		salary: 2000,
		pf:     20,
	}
	p3 := PermanentEmployee{
		ID:     3,
		salary: 3000,
		pf:     30,
	}
	t1 := TemporaryEmployee{
		ID:     4,
		salary: 800,
	}
	f1 := Freelancer{
		ID:          5,
		ratePerHour: 20,
		totalHours:  80,
	}
	employees := []SalaryCalculator{p1, p2, p3, t1, f1}
	totalExpenses(employees)
}
