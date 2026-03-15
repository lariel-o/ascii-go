package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"math"
)

func main() {
	currentTableType := "printable"

	if(len(os.Args) > 1) {
		currentTableType = os.Args[1]
	}

	currentTableMin 		:= tableType[currentTableType][0]
	currentTableMax 		:= tableType[currentTableType][1]
	currentTableLength		:= currentTableMax - currentTableMin 

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.AlignRight)
	header := "Dec\tHex\tSymb\t\tDec\tHex\tSymb\t\tDec\tHex\tSymb\t\tDec\tHex\tSymb\t\tDec\tHex\tSymb\t\tDec\tHex\tSymb\t\t"
	fmt.Fprintln(w, header)

	loopCount := int(math.Round(float64(currentTableLength) / 6))
	baseLine := currentTableMin
	for i := range loopCount {
		if i != 0 && i % 8 == 0 {
			w.Flush()
			fmt.Fprintln(w, "\n" + header)
			baseLine += 41
		}

		runningLine := baseLine + i 
		line := ""


		for j := range 6  {
			currentDec := runningLine + j*8
			currentHex := fmt.Sprintf("%x", currentDec)
			currentSym := string(currentDec)

			if currentTableMax < currentDec {
				break
			}

			line += fmt.Sprintf(" %d\t%s\t%s\t\t", currentDec, currentHex, currentSym)
		}

		fmt.Fprintln(w, line)
	}

	// talvez (?) precisa de um metodo para verificar se loopCount/8 é um inteiro
	w.Flush()
}

