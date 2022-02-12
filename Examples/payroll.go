// Developer: Eddie Rangel
// File name: payroll.go
// Description: Calculates an employees gross pay
// Copyright (c) 2022

package main

import "fmt"

func main() {

	fmt.Println("\n")
	fmt.Println("Welcome to Our Payroll Manager!!\n")

	// Gross Pay is Hours Worked multiplied by Hourly Payrate

	var gross_pay, hours_worked, pay_rate int

	pay_rate = 20
	hours_worked = 40
	gross_pay = pay_rate * hours_worked

	fmt.Println("Week One Payroll\n")
	fmt.Println("Employee\tGross Pay\tPay Rate\tHours Worked")
	fmt.Println("_____________________________________________________")

	fmt.Printf("F. Name\t\t$%d\t\t$%d\t\t%d\n", gross_pay, pay_rate, hours_worked)
	fmt.Printf("F. Name\t\t$%d\t\t$%d\t\t%d\n", gross_pay, pay_rate*2, hours_worked)
	fmt.Printf("F. Name\t\t$%d\t\t$%d\t\t%d\n", gross_pay, pay_rate*3, hours_worked)
	fmt.Println("\n")
}
