package dts

import (
	"fmt"
)

func MultipleDeclare() {
	j, k, l := "shark", 2.05, 15
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}

const (
	year     = 365 // When the program compiles, it explicitly converts years to an int
	leapYear = int32(366)
)

func ConstVars() {
	hours := 24
	minutes := int32(60)
	fmt.Println(hours * year)
	fmt.Println(minutes * year)
	fmt.Println(minutes * leapYear)

    // INVALID
    // fmt.Println(hours * leapYear)
}
