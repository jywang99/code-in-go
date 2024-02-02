package main

import "fmt"

var maxLvl = 3
func printHeader(title string, lvl int) {
    head := "\n"
    for i := 0; i < maxLvl-lvl+1; i++ {
        head += "#"
    }
    fmt.Printf("%s %s\n", head, title)
}

