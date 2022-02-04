package main

// File name: arguments.go
// Description: Displays command line arguments

import (
	"fmt"
	"os"
)

func main() {

	var element, seperator string

	//fmt.Println("Filename: ", os.Args[0])
	fmt.Println("Arguments Assignment\n")

	for index := 1; index < len(os.Args); index++ {
		element += seperator + os.Args[index]
		seperator = " "
	}
	fmt.Println(element)

}
