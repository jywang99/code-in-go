package flows

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func BuildErrors() {
	err := errors.New("barnacles")
	fmt.Println("Sammy says:", err)
	err2 := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", err2)
}

func boom() error {
	return errors.New("barnacles")
}

func CatchFnError() {
	err := boom()
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("Anchors away!")
}

func capitalize(name string) (string, int, error) {
	handle := func(err error) (string, int, error) {
		return "", 0, err
	}
	if name == "" {
		return handle(errors.New("no name provided"))
	}
	return strings.ToTitle(name), len(name), nil
}

func CatchCapitalizeError() {
	name, size, err := capitalize("sammy")
	if err != nil {
		fmt.Println("An error occurred:", err)
	}
	fmt.Printf("Capitalized name: %s, length: %d\n", name, size)
}
