package flows

import (
	"errors"
	"fmt"
	"net/http"
	// "os"
	"log"
)

type RequestError struct {
	StatusCode int
	Err        error
}

// func (r *RequestError) Error() string {
//     return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
// }

func doRequest() error {
	return &RequestError{
		StatusCode: 503,
		Err:        errors.New("unavailable"),
	}
}

func CatchCustomError() {
	err := doRequest()
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
		return
	}
	fmt.Println("success!")
}

/* Custom methods for errors */

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func (r *RequestError) Temporary() bool {
	return r.StatusCode == http.StatusServiceUnavailable // 503
}

func HandleWithRetryIfTemporary() {
	err := doRequest()
	if err != nil {
		fmt.Println(err)
		re, ok := err.(*RequestError) // type assertion, re: RequestError
		if ok {
			if re.Temporary() {
				fmt.Println("This request can be tried again")
			} else {
				fmt.Println("This request cannot be tried again")
			}
		}
		// os.Exit(1)
		return
	}
	fmt.Println("success!")
}

/* Wrapping errors */

type WrappedError struct {
	Context string
	Err     error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

func Wrap(err error, info string) *WrappedError {
	return &WrappedError{
		Context: info,
		Err:     err,
	}
}

func TestWrappedError() {
	err := errors.New("boom!")
	err = Wrap(err, "TestWrappedError")
	fmt.Println(err)
}

/* Recover from panic */

func RecoverFromPanic() {
	divideByZero()
	fmt.Println("we survived dividing by zero!")
}

func divideByZero() {
	defer func() {
		if err := recover(); err != nil { // recover() stops panicing and returns error
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
	return a / b
}
