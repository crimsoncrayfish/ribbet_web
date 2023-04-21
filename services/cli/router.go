package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"ribbit_web/activity"
	inputhandler "ribbit_web/core/input_handler"
	"ribbit_web/core/twx"
	typeconvert "ribbit_web/core/type_convert"
	"strings"
	"text/tabwriter"
)

type CliRouter struct {
	activityStore activity.ActivityFileStore
	writer        *tabwriter.Writer
	logger        *log.Logger
	inputReader   *bufio.Reader
	inputForm     inputhandler.FormController
}

func InitialiseRouter(activityStore activity.ActivityFileStore, l *log.Logger, reader *bufio.Reader) CliRouter {
	tw := twx.New(os.Stdout)
	return CliRouter{
		activityStore: activityStore,
		writer:        tw,
		logger:        l,
		inputReader:   reader,
		inputForm:     inputhandler.InitialiseController(),
	}
}

func (r CliRouter) HandleInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	instruction, args, ok := strings.Cut(input, " ")
	if ok {
		if err := r.HandleInstructionWithArgs(instruction, args); err != nil {
			return err
		}
	} else {
		if err := r.HandleInstruction(instruction); err != nil {
			return err
		}
	}
	return nil
}

func (r CliRouter) HandleInstruction(instruction string) error {
	switch instruction {
	case "l":
		r.logger.Println("INFO: Listing all activities...")
		r.logger.Println()
		list, err := r.activityStore.List()
		if err != nil {
			return err
		}
		r.displayActivities(list)
	case "d":
		r.logger.Println("INFO: No arguments found for delete command. Requesting activity id to delete...")
		id, err := r.inputForm.GetString("ID to DELETE")
		if err != nil {
			return err
		}
		if err := r.activityStore.Delete(id); err != nil {
			if activity.IsNotFoundError(err) {
				r.logger.Printf("WARNING: Could not find activity with id %s to delete...", id)
				return nil
			}
			return err
		}
	case "c":
		r.logger.Println("INFO: No arguments found for create command. Requesting activity details...")

		user, err := r.inputForm.GetString("User Name")
		if err != nil {
			return err
		}
		description, err := r.inputForm.GetString("Description")
		if err != nil {
			return err
		}
		activityType, err := r.inputForm.GetActivityType("Activity Type")
		if err != nil {
			return err
		}
		date, err := r.inputForm.GetDate("Activity Date")
		if err != nil {
			return err
		}

		if err := r.activityStore.Create(user, description, activityType, date); err != nil {
			return err
		}
	case "q":
		return activity.ErrCloseApplication
	case "h":
		r.logger.Print(routerControls)
	default:
		r.logger.Printf("WARNING: Unknown command %s", instruction)
		r.logger.Println([]byte("l"))
		r.logger.Println([]byte(instruction))
	}
	return nil
}

func (r CliRouter) HandleInstructionWithArgs(instruction string, args string) error {
	switch instruction {
	case "l":
		r.logger.Println("INFO: List command takes no arguments...")
	case "d":
		if err := r.activityStore.Delete(args); err != nil {
			if activity.IsNotFoundError(err) {
				r.logger.Printf("WARNING: Could not find activity with id %s to delete...", args)
				return nil
			}
			return err
		}
	case "c":
		r.logger.Println("Creating new activity")
		arguments := strings.Split(args, ",")
		if len(arguments) != 4 {
			r.logger.Println("WARNING: Create command requires exactly 4 arguments seperated by ','. Please try again...")
		} else {
			user := arguments[0]
			description := arguments[1]
			activityType, err := activity.ToType(arguments[2])
			if err != nil {
				return err
			}
			date, err := typeconvert.StringToDate(arguments[3])
			if err != nil {
				return err
			}
			if err := r.activityStore.Create(user, description, activityType, date); err != nil {
				return err
			}
		}
	case "q":
		r.logger.Println("WARNING: Quit command takes no arguments...")
	case "h":
		r.logger.Print("WARNING: Help command takes no arguments...")
	default:
		r.logger.Printf("WARNING: Unknown command %s", instruction)
		r.logger.Println([]byte("l"))
		r.logger.Println([]byte(instruction))
	}

	return nil
}

func (r CliRouter) displayActivities(list []activity.Activity) {
	//use tabwriter to format the output
	fmt.Fprintf(r.writer, "%s\t%s\t%s\t%s\t%s\n", "Id", "Username", "Description", "Type", "Date completed")
	for _, a := range list {
		fmt.Fprintf(r.writer, "%s\t%s\t%s\t%s\t%s\n", a.Id, a.UserName, a.Description, a.ActivityType.ToString(), a.CompletedAt)
	}

	r.writer.Flush()
	fmt.Println()
}

const routerControls string = `Ribbit Habbit Tracker:

Usage:
l                                                    - list all activities
d [ID]                                               - delete Activity with ID [ID]
c [User,Description,ActivityType,CompletionTime]     - create new Activity with details [User,Description,ActivityType,CompletionTime]
u [ID,User,Description,ActivityType,CompletionTime]  - update Activity with details [ID,User,Description,ActivityType,CompletionTime]
q                                                    - quit TodoList
h                                                    - display this menu again
--------------------------------------------------
`
