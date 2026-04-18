package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the program... Goodbye!")
	os.Exit(0)
	return nil
}
