package strs

import (
    "fmt"
    "strings"
)

func HelloInput() {
	fmt.Println("Enter your name.")
	var name string
	// fmt.Scanln(&name)
    name = "jy "
	name = strings.TrimSpace(name) // not required in go 1.21.6 to trim trailinig newline
	fmt.Printf("Hi, %s! I'm Go!\n", name)
}

func Utf8chars() {
	a := "Hello, 世界"
	for i, c := range a {
		fmt.Printf("%d: %s\n", i, string(c))
	}
	fmt.Println("length of 'Hello, 世界': ", len(a)) // 13 because multi-byet chars
	// 世界 are runes, not int32s
}

