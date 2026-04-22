package main

import "fmt"

// 📂 05_oop
// Proj: Payroll-Processor

// Our custom interface
type Payable interface{
	CalculatePay() float64
}

type SalariedEmployee struct {
	Name string
	AnnualSalary float64
}

func (se SalariedEmployee) CalculatePay()float64{
	return se.AnnualSalary/12.0
}

// Stringer interface{ }
func (se SalariedEmployee) String()string{
	return  fmt.Sprintf("Salaried: %s (Annual: $%.2f)",se.Name,se.AnnualSalary)
}

type HourlyEmployee struct {
	Name string
	HourlyRate float64
	HourseWorked float64 // Hrs worked in the month
}

func (he HourlyEmployee) CalculatePay()float64{
	return he.HourlyRate*he.HourseWorked
}

// Stringer interface{ }
func (he HourlyEmployee) String()string{
	return  fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %1.f)",he.Name,he.HourlyRate,he.HourseWorked)
}

type CommissionedEmployee struct{
		Name string
		BaseSalary float64 // Monthly base
		CommissionRate float64 // e.g., 0.05 for 5%
		SalesAmount float64

}

func (ce CommissionedEmployee) CalculatePay()float64{
	return ce.BaseSalary + (ce.CommissionRate * ce.SalesAmount)
}

// Stringer interface{ }
func (ce CommissionedEmployee) String()string{
	return  fmt.Sprintf("Commission: %s (Base: $%.2f, CommRate: %.2f%%, Sales:   $%.2f)", ce.Name, ce.CommissionRate*100, ce.SalesAmount)
}


func main() {
	
}