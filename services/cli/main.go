package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"ribbit_web/activity"
	filestore "ribbit_web/core/file_store"
	"ribbit_web/core/logx"
)

func main() {
	verbose := flag.Bool("verbose", false, "pass the verbose flag to log output to the cli")
	flag.Parse()

	l := logx.New("Ribbit CLI", *verbose)

	if err := run(l); err != nil {
		if activity.IsCloseApplicationError(err) {
			l.Println("Shutting down application...")
			os.Exit(0)
		}
		fmt.Println("Unexpected Error, closing application")
		l.Printf("ERROR: An unexpected Error occurred '%s", err)
		os.Exit(1)
	}
}

func run(l *log.Logger) error {
	//create filestore file
	databaseFile, err := filestore.InitDatabaseFile(l, "FILE_STORE_LOCATION")
	if err != nil {
		return err
	}
	defer func() {
		l.Println("Start closing database file...")
		if closeErr := databaseFile.Close(); closeErr != nil {
			l.Printf("ERROR: closing file: %s", err)
		} else {
			l.Println("Database file close successfull...")
		}
	}()

	//initialise db store and input reader
	l.Println("Initialising...")
	store := activity.NewActivityFileStore(databaseFile, l)
	router := InitialiseRouter(store, l)

	//listen for inputs
	l.Println("Creating input reader...")
	reader := bufio.NewReader(os.Stdin)
	l.Println("--------------Initialisation complete-------------")
	l.Print(routerControls)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		if err = router.HandleInput(line); err != nil {
			return err
		}
		l.Println()
		l.Println("---------------------------")
	}
}
