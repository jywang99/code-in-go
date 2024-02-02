package strs

import (
    "fmt"
    "strings"
)

func StringsLib() {
    ss := "Sammy Shark"
    fmt.Println(strings.HasPrefix(ss, "Sammy"))
    fmt.Println(strings.HasSuffix(ss, "Shark"))
    fmt.Println(strings.Contains(ss, "Sh"))

    // join array to string
    fmt.Println(strings.Join([]string{"sharks", "crustaceans", "plankton"}, ","))

    // split string to array
    balloon := "Sammy has a balloon."
    s := strings.Split(balloon, " ")
    fmt.Println(s)

    // replace
    fmt.Println(strings.ReplaceAll(balloon, "has", "had"))
}
