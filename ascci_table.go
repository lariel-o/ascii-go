package main

import (
	"fmt"
	// "github.com/olekukonko/tablewriter"
)



func getASCIIRelation(interval[2] int, extended bool, detailed bool) [][3]string { // mudar 3->4
	toReturn := make([][3]string, interval[1]-interval[0]+1) // mudar 3->4

	for i := range len(toReturn) {
		currentDecimalValue := interval[0] + i
		toReturn[i] = [3]string{
			fmt.Sprintf("%d", 		currentDecimalValue), 		// decimal
			fmt.Sprintf("%2x", 		currentDecimalValue), 		// hexadecimal
			fmt.Sprintf("%s", string(currentDecimalValue)),		// symbol
		} //mudar 3->4
	}

	return toReturn
}

func main() {
	getASCIIRelation([2]int{32, 127}, false)
	fmt.Println("running")	
}

