package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
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

	// choose if will render 8 elements per table or 16 per table
	separator := 8
	if slices.Contains(os.Args, "-e") {
		separator = 16
	}

	interval := tables_type["printable"]

	// choose interval by max
	ind := slices.Index(os.Args[1:], "--max")
	if ind != -1 {
		v, _ := strconv.Atoi(os.Args[ind+2])
		interval[1] = v
	}

	// choose interval by min
	ind = slices.Index(os.Args[1:], "--min")
	if ind != -1 {
		v, _ := strconv.Atoi(os.Args[ind+2])
		interval[0] = v
	}

	// choose interval by table type
	ind = slices.Index(os.Args[1:], "-t")
	if ind != -1 {
		interval = tables_type[os.Args[ind+2]]
	}

	data := getASCIIRelation(interval)
	for i := range len(data)/separator {
		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"Dec", "Hex", "Symbol", "Description"})
		table.Bulk(data[i*separator : separator*(i+1)])
		table.Render()
	}
}

