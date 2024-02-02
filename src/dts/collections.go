package dts

import (
	"fmt"
)

func PrintCollections() {
	// array (fixed length)
	coral := [3]string{"blue coral", "staghorn coral", "pillar coral"}
	fmt.Println(coral)
	// slice (dynamic length)
	seaCreatures := []string{"shark", "cuttlefish", "squid", "mantis shrimp"}
	fmt.Println(seaCreatures)
	// map
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue"}
	fmt.Println(sammy)
	fmt.Println(sammy["color"])
}

func IterateMaps() {
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue"}
	fmt.Println(sammy)

	// iterate key, val
	for key, value := range sammy {
		fmt.Printf("%q is the key for the value %q\n", key, value)
	}
	// iterate key only
	keys := []string{}
	for key := range sammy {
		keys = append(keys, key)
	}
	fmt.Printf("keys: %q\n", keys)
	// values only
	var i int
	items := make([]string, len(sammy))
	for _, v := range sammy {
		items[i] = v
		i++
	}
	fmt.Printf("values: %q\n", items)

	// intialized with zero values
	counts := map[string]int{}
	fmt.Println(counts["sammy"]) // 0
	names := map[string]string{}
	fmt.Println(names["sammy"]) // empty string

	// key existence checking
	count, ok := counts["sammy"]
	if ok {
		fmt.Printf("Sammy has a count of %d\n", count)
	} else {
		fmt.Println("Sammy was not found")
	}

    // inline
    counts["sammy"] = 1
    if count, ok := counts["sammy"]; ok {
        fmt.Printf("Sammy has a count of %d\n", count)
    } else {
        fmt.Println("Sammy was not found")
    }
}

func ArraysAndSlices() {
    coral := [4]string{"blue coral", "foliose coral", "pillar coral", "elkhorn cor"}
    fmt.Printf("%q\n", coral)
    // append to array = error

    fmt.Println("Append black coral")
    coralSlice := coral[:] // slice with all elements
    coralSlice = append(coralSlice, "black coral") // length is now variable
    fmt.Printf("%q\n", coralSlice)
    // append multiple
    fmt.Println(`Append "massive coral", "soft coral"`)
    coralSlice = append(coralSlice, "massive coral", "soft coral") // length is now variable
    fmt.Printf("%q\n", coralSlice)

    // concatenate slices
    fmt.Println(`Concat "massive coral", "soft coral"`)
    moreCoral := []string{"massive coral", "soft coral"}
    coralSlice = append(coralSlice, moreCoral...)

    // remove item
    fmt.Println("Remove index 3")
    coralSlice = append(coralSlice[:3], coralSlice[4:]...)
    fmt.Printf("%q\n", coralSlice)

    // multidimensional slice
    seaNames := [][]string{{"shark", "octopus", "squid", "mantis shrimp"}, {"Sammy", "Jessie", "Drew", "Jamie"}}
    fmt.Printf("Multidimensional slice: %q\n", seaNames)
    fmt.Printf("seaNames[1][0]: %q\n", seaNames[1][0])
    fmt.Printf("seaNames[0][0]: %q\n", seaNames[0][0])
}
