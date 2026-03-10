package main

var tables_type = map[string][2]int{
	"printable": 		[2]int{32, 127},
	"control":   		[2]int{0, 31},
	"extended": 		[2]int{128, 255},
	"letters": 			[2]int{65, 122},
}

