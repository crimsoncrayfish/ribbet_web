package inputhandler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FormController struct {
	ioHandler *bufio.Reader
}

func InitialiseController() FormController {
	return FormController{
		ioHandler: bufio.NewReader(os.Stdin),
	}
}

func (controller FormController) GetValuesFor(inputs ...string) (map[string]string, error) {

	fmt.Printf("Initializing form for %v...\n", inputs)

	output := make(map[string]string)
	for i := 0; i < len(inputs); i++ {
		input := strings.TrimSpace(inputs[i])
		value, err := controller.getValuesForInput(input)
		if err != nil {
			return nil, err
		}
		output[input] = value
	}
	fmt.Println("\nForm data collected...")
	return output, nil
}

func (controller FormController) getValuesForInput(argName string) (string, error) {
	fmt.Print("Please enter " + argName + ": ")
	line, err := controller.ioHandler.ReadString('\n')
	if err != nil {
		return "", err
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	return line, nil
}
