package inputhandler

import (
	"bufio"
	"fmt"
	"os"
	"ribbit_web/activity"
	typeconvert "ribbit_web/core/type_convert"
	"strings"
	"time"
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
	for index := range inputs {
		value, err := controller.GetString(inputs[index])
		if err != nil {
			return nil, err
		}
		output[inputs[index]] = value
	}
	fmt.Println("\nForm data collected...")
	return output, nil
}

func (controller FormController) GetString(argumentName string) (string, error) {
	fmt.Print("Please enter " + argumentName + ": ")
	line, err := controller.ioHandler.ReadString('\n')
	if err != nil {
		return "", err
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")

	return line, nil
}

func (controller FormController) GetActivityType(argumentName string) (activity.Type, error) {
	line, err := controller.GetString(argumentName)
	if err != nil {
		return activity.None, err
	}

	activityType, err := activity.ToType(line)
	if err != nil {
		fmt.Printf("WARNING: Failed to interpret '%s'. Expected values are %v\n", line, activity.ListTypeOptions())
		return controller.GetActivityType(argumentName)
	}

	return activityType, nil
}

func (controller FormController) GetDate(argumentName string) (time.Time, error) {
	line, err := controller.GetString(argumentName)
	if err != nil {
		return time.Now(), err
	}

	date, err := typeconvert.StringToDate(line)
	if err != nil {
		return controller.GetDate(argumentName)
	}

	return date, nil
}

/*
	out, err := castToType(line, argumentType)
	if err != nil {
		if !(IsTypeNotFoundError(err) || typeconvert.IsStringToDateConvertFailedError(err)) {
			return "", err
		}
		fmt.Printf("ERROR: Could not convert input '%s' to type '%s'\n", line, argumentType)
		return controller.getValueForInput(argumentName, argumentType)
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
	case DateArg:
		return typeconvert.StringToDate(input)
	default:
		return "", ErrTypeNotFound
	}
}*/
