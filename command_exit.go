package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Exiting the app...")
	os.Exit(0)

	return errors.New("app is still running")
}
