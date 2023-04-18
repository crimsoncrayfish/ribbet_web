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

func run(logger *log.Logger) error {
	//create filestore file
	databaseFile, err := filestore.InitDatabaseFile(logger, "FILE_STORE_LOCATION")
	if err != nil {
		return err
	}
	defer func() {
		logger.Println("Start closing database file...")
		if closeErr := databaseFile.Close(); closeErr != nil {
			logger.Printf("ERROR: closing file: %s", err)
		} else {
			logger.Println("Database file close successfull...")
		}
	}()

	//initialise db store
	logger.Println("Initialising...")
	store := activity.NewActivityFileStore(databaseFile, logger)

	//initialise input reader
	logger.Println("Creating input reader...")
	reader := bufio.NewReader(os.Stdin)

	//listen for inputs
	logger.Println("Creating router...")
	router := InitialiseRouter(store, logger, reader)
	logger.Println("--------------Initialisation complete-------------")
	logger.Print(routerControls)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		if err = router.HandleInput(line); err != nil {
			return err
		}
		logger.Println()
		logger.Println("---------------------------")
	}
}
