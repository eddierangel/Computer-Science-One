// Developer: Eddie Rangel
// File name: arguments.go
// Description: Displays command line arguments
// Copyright (c) 2022

package main

import (
	"fmt"
	"os"
)

func main() {

	var element, seperator string

	//fmt.Println("Filename: ", os.Args[0])
	fmt.Println("\n")
	fmt.Println("Process Command Line Arguements...\n")

	for index := 1; index < len(os.Args); index++ {
		element += seperator + os.Args[index]
		seperator = " "
	}
	fmt.Println(element)
	fmt.Println("\n")

}
