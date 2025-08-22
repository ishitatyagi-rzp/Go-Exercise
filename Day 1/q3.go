package main

import "fmt"

type emp interface {
	CalcSalary() int
}

type FullTime struct {
	days      int
	payPerDay int
}

func (f FullTime) CalcSalary() int {
	return f.days * f.payPerDay
}

type Contractor struct {
	days      int
	payPerDay int
}

func (c Contractor) CalcSalary() int {
	return c.days * c.payPerDay
}

type Freelancer struct {
	hours      int
	payPerHour int
}

func (f Freelancer) CalcSalary() int {
	return f.hours * f.payPerHour
}

func main() {
	ft := FullTime{days: 30, payPerDay: 500}
	ct := Contractor{days: 10, payPerDay: 300}
	fr := Freelancer{hours: 20, payPerHour: 100}
	employees := []emp{ft, ct, fr}
	for i, e := range employees {
		fmt.Printf("Employee %d salary: %d\n", i+1, e.CalcSalary())
	}
}
