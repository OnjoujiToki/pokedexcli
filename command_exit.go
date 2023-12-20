package main

import (
	"fmt"
	"os"
)

func printExit(defaultConfig *configClient) error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
