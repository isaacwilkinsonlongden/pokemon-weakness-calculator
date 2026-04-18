package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, words ...string) error {
	fmt.Println("Closing the program... Goodbye!")
	os.Exit(0)
	return nil
}
