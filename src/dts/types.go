package dts

import (
    "fmt"
    "log"
    "strconv"
)

func NumberToString() {
	lines_yesterday := "50"
	lines_today := "108"
	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil { // error handling
		log.Fatal(err) // log and exit program
	}
	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
	}
	lines_more := today - yesterday
	fmt.Println(lines_more)
}

func StringToBytes() {
    a := "my string"
    b := []byte(a)
    c := string(b)
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}

func FloatAddition() {
    e := 5.5
    f := 2.5
    s := e + f
    fmt.Println(s)
    fmt.Printf("%.2f\n", s)

    i := 7
    j := 7.0
    // invalid operation: i + j (mismatched types int and float64) [MismatchedTypes]
    // fmt.Println(i + j)
    fmt.Printf("%.2f\n", float64(i) + j) // cast works
}

func Division() {
    m := 80
    n := 6
    fmt.Printf("m=%d, n=%d\n", m, n)
    fmt.Printf("m / n: %d\n", m/n)
    fmt.Printf("float64(m) / float64(n): %.2f\n", float64(m) / float64(n))

}
