// Developer: Eddie Rangel
// File name: FahrenheitToCelsius.go
// Description: Displays a table of Fahrenheit and Celsius values
// Copyright (c) 2022

package main

import "fmt"

func main() {

	fmt.Println("\n")
	fmt.Println("Display Fahrenheit and Celsius values...\n")

	var fahrenheit, celsius int
	var upper, step int

	fahrenheit = 0
	upper = 300
	step = 20

	fmt.Println("Fahrenheit\tCelsius")
	fmt.Println("_______________________")

	for index := 0; index <= upper; index += step {
		celsius = 5 * (fahrenheit - 32) / 9
		fmt.Printf("%d\t\t%d\n", fahrenheit, celsius)
		fahrenheit += step
	}
	fmt.Println("\n")

}
