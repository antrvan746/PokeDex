package main

import (
	"fmt"
	"os"
	"errors"
)

func commandExit(cfg *config) error {
	fmt.Println("Exiting the app...")
	os.Exit(0)

	return errors.New("app is still running")
}