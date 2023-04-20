package inputhandler

import (
	"bufio"
	"fmt"
	"net/mail"
	"os"
	"strconv"
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

func (controller FormController) GetValuesFor(inputs ...string) (map[string]any, error) {

	fmt.Printf("Initializing form for %v...\n", inputs)
	argsMap, err := getAllArgumentsAndTypes(inputs...)
	if err != nil {
		return nil, err
	}

	output := make(map[string]any)
	for k, v := range argsMap {
		value, err := controller.getValuesForInput(k, v)
		if err != nil {
			return nil, err
		}
		output[k] = value
	}
	fmt.Println("\nForm data collected...")
	return output, nil
}

func getAllArgumentsAndTypes(inputs ...string) (map[string]string, error) {
	splitInputs := make(map[string]string)
	for i := 0; i < len(inputs); i++ {
		argumentName, argumentType, err := getArgumentAndType(inputs[i])
		if err != nil {
			return nil, err
		}
		splitInputs[argumentName] = argumentType
	}
	return splitInputs, nil
}

func getArgumentAndType(input string) (string, string, error) {
	argumentName, argumentType, ok := strings.Cut(input, ":")
	if !ok {
		argumentType = "string"
	} else {
		if _, err := isArgumentTypeValid(argumentType); err != nil {
			return "", "", err
		}
	}
	return argumentName, argumentType, nil
}

func (controller FormController) getValuesForInput(argumentName, argumentType string) (any, error) {
	fmt.Print("Please enter " + argumentName + ": ")
	line, err := controller.ioHandler.ReadString('\n')
	if err != nil {
		return "", err
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")

	out, err := castToType(line, argumentType)
	if err != nil {
		return "", err
	}
	return out, nil
}

var ArgumentTypes = map[string]struct{}{StringArg: {}, IntegerArg: {}, DoubleArg: {}, DateArg: {}, EmailArg: {}, PhoneNumberArg: {}, ActivityTypeArg: {}}

const (
	StringArg       = "string"
	IntegerArg      = "integer"
	DoubleArg       = "double"
	DateArg         = "date"
	EmailArg        = "email"
	PhoneNumberArg  = "phone_number"
	ActivityTypeArg = "activity_type"
)

func isArgumentTypeValid(argumentTypeString string) (bool, error) {
	_, found := ArgumentTypes[strings.ToLower(argumentTypeString)]
	if !found {
		return found, ErrTypeNotFound
	}
	return found, nil
}

func castToType(input, argumentType string) (any, error) {
	switch argumentType {
	case StringArg:
		return input, nil
	case IntegerArg:
		return strconv.Atoi(input)
	case DoubleArg:
		val, err := strconv.ParseFloat(input, 32)
		if err != nil {
			val, err = strconv.ParseFloat(input, 64)
		}
		return val, err
	case EmailArg:
		_, err := mail.ParseAddress(input) //validate as email adress
		return input, err
	default:
		return "", ErrTypeNotFound
	}
}
