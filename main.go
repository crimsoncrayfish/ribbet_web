package main

import (
	"log"
	"os"
	logr "ribbet_web/common/log_r"
)

// TASKS:
//1. File store db for activities {id, description, time}
//2. CRUD for activities
//3. Basic Activities ui

func main() {
	l := logr.New("My Fancy Logger")

	if err := run(l); err != nil {
		l.Printf("Failed with error %s", err)
		os.Exit(1)
	}
}

func run(l *log.Logger) error {
	l.Println("Running Application")

	return nil
}
