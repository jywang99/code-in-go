package buildtags

import "fmt"

var features = []string{
    "Free Feature #1",
    "Free Feature #2",
}

func PrintFeatures() {
    for _, f := range features {
        fmt.Println(">", f)
    }
}

// To add pro features: run with
// go run -tags pro src/*.go 

// To add pro and enterprise features: run with
// go run -tags pro,enterprise src/*.go 
