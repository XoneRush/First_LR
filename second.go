package main

import (
	"fmt"
)

func second() {
	var (
		intVar    int     = 1
		floatVar  float64 = 5.2
		boolVar   bool    = true
		stringVar string  = "this is string var"
	)

	fmt.Println("\nintVar: ", intVar)
	fmt.Println("floatVar: ", floatVar)
	fmt.Println("boolVarL ", boolVar)
	fmt.Println("stringVar: ", stringVar)
}
