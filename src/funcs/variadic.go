package funcs

import "fmt"

func Example() {
	fmt.Println()
	fmt.Println("one")
	fmt.Println("one", "two")
	fmt.Println("one", "two", "three")
}

/* defining variadic function */
func sayHello(names ...string) {
	if len(names) == 0 {
		fmt.Println("nobody to greet")
		return
	}
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}

func join(del string, values ...string) string { // varadic arg must be the last one!
	var line string
	for i, v := range values {
		line = line + v
		if i != len(values)-1 {
			line = line + del
		}
	}
	return line
}

func VariadicFunc() {
	sayHello()
	sayHello("Sammy")
	sayHello("Sammy", "Jessica", "Drew", "Jamie")

	fmt.Println(join(", ", "Sammy", "Jessica", "Drew", "Jamie"))
	fmt.Println(join(", ", "Sammy"))
	fmt.Println(join(", "))

    // explode
    names := []string{"Amy", "Bianca", "Cassy"}
    sayHello(names...)
}
