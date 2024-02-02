package flows

import (
	"io"
	"log"
	"os"
)

func TestDefer() {
	if err := write("readme.txt", "This is a readme file"); err != nil {
		log.Fatal("failed to write file:", err)
	}
}

// PRORBLEM: if error occurs before the defer, file is not closed
func write(fileName string, text string) error {
	file, err := os.Create(fileName)
    if err != nil {
		return err
	}
	defer file.Close() // finally
	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}
	return file.Close() // so that any closing error is returned
}
