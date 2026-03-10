package main

import (
	"fmt"
	"os"
	"slices"
	"github.com/olekukonko/tablewriter"
)

func getASCIIRelation(interval[2] int) [][]string { 
	toReturn := make([][]string, interval[1]-interval[0]+1) 

	for i := range len(toReturn) {
		currentDecimalValue := interval[0] + i
		toReturn[i] = []string{
			fmt.Sprintf("%d", 		currentDecimalValue), 				// decimal
			fmt.Sprintf("%2x", 		currentDecimalValue), 				// hexadecimal
			fmt.Sprintf("%s", string(currentDecimalValue)),				// symbol
			fmt.Sprintf("%s", asciiDescription[currentDecimalValue]),	// description
		} 
	}

	return toReturn
}

func main() {
	data := getASCIIRelation([2]int{32, 127})

	separator := 8
	if slices.Contains(os.Args, "-e") {
		separator = 16
	}

	for i := range len(data)/separator {
		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"Dec", "Hex", "Symbol", "Description"})
		table.Bulk(data[i*separator : separator*(i+1)])
		table.Render()
	}
}

