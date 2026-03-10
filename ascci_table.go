package main

import (
	"fmt"
	// "github.com/olekukonko/tablewriter"
)

func getASCIIRelation(interval[2] int) [][4]string { 
	toReturn := make([][4]string, interval[1]-interval[0]+1) 

	for i := range len(toReturn) {
		currentDecimalValue := interval[0] + i
		toReturn[i] = [4]string{
			fmt.Sprintf("%d", 		currentDecimalValue), 		// decimal
			fmt.Sprintf("%2x", 		currentDecimalValue), 		// hexadecimal
			fmt.Sprintf("%s", string(currentDecimalValue)),		// symbol
			fmt.Sprintf("%s", asciiDescription[currentDecimalValue]),
		} 
	}

	return toReturn
}

func main() {
	fmt.Println(getASCIIRelation([2]int{32, 127}))
	fmt.Println("running")	
}

