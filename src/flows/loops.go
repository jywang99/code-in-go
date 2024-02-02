package flows

import (
	"bytes"
	"fmt"
	"io"
)

func ReadBytesLoop() {
	buf := bytes.NewBufferString("one\ntwo\nthree\nfour\n")
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print(line)
				break
			}
			fmt.Println(err)
			break
		}
		fmt.Print(line)
	}
}

/* range clause */
func RangeClause() {
    // iterate
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}
	for i, shark := range sharks {
		fmt.Println(i, shark)
	}

    // iterate through chars in string
    str := "Hello world"
    fmt.Printf(" | ")
    for _, c := range(str) {
        fmt.Printf("%c | ", c)
    }
    fmt.Println()

    // fill an array
    integers := make([]int, 10)
    fmt.Println(integers)
    for i := range integers {
        integers[i] = i
    }
    fmt.Println(integers)
}
