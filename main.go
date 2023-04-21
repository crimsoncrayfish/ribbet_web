package main

import (
	"bufio"
	"fmt"
	"os"
	inputhandler "ribbit_web/core/input_handler"
	"strings"
)

func main() {
	fmt.Println("Testing something...")
	if err := run(); err != nil {
		fmt.Printf("Unexpected error occurred. Err: %s", err)
		os.Exit(1)
	}

}

func run() error {

	//initialise input reader
	fmt.Println("Creating input reader...")
	handler := inputhandler.InitialiseController()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please provide a comma seperated list of argument names: ")
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")
		if err != nil {
			return err
		}

		arguments := strings.Split(line, ",")
		output, err := handler.GetValuesFor(arguments...)
		if err != nil {
			return err
		}
		printMap(output)
		fmt.Println("-------------------------------------------------")

	}
}

func printMap(myMap map[string]string) {
	for k, v := range myMap {
		fmt.Printf("%s: %s\n", k, v)
	}
}
