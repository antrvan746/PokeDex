package main

import (
	"fmt"
	"os"
	"errors"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Exiting the app...")
	os.Exit(0)

	return errors.New("app is still running")
}