package main

import (
	"fmt"
	"log"
	"os"
	"ribbit_web/activity"
	"ribbit_web/core/twx"
	"strings"
	"text/tabwriter"
	"time"
)

type CliRouter struct {
	activityStore activity.ActivityFileStore
	writer        *tabwriter.Writer
	logger        *log.Logger
}

func InitialiseRouter(activityStore activity.ActivityFileStore, l *log.Logger) CliRouter {
	tw := twx.New(os.Stdout)
	return CliRouter{
		activityStore: activityStore,
		writer:        tw,
		logger:        l,
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
		r.logger.Println("Listing all activities...")
		r.logger.Println()
		list, err := r.activityStore.List()
		if err != nil {
			return err
		}
		r.displayActivities(list)
	case "d":
		r.logger.Println("INFO: Delete command requires 1 argument...")
	case "c":
		r.logger.Println("INFO: Create command requires 4 arguments...")
	case "q":
		return activity.ErrCloseApplication
	case "h":
		r.logger.Print(routerControls)
	default:
		r.logger.Printf("Unknown command %s", instruction)
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
		r.logger.Printf("Deleting activity with id %s", args)
		if err := r.activityStore.Delete(args); err != nil {
			if activity.IsNotFoundError(err) {
				r.logger.Printf("WARN: Could not find activity with id %s to delete...", args)
				return nil
			}
			return err
		}
	case "c":
		r.logger.Println("Creating new activity")
		arguments := strings.Split(args, ",")
		if len(arguments) != 4 {
			r.logger.Println("WARN: Create command requires exactly 4 arguments...")
		} else {
			user := arguments[0]
			description := arguments[1]
			activityType, err := activity.ToType(arguments[2])
			if err != nil {
				if activity.IsActivityTypeNotFoundError(err) {
					r.logger.Printf("WARN: Ensure that activity type format is correct. Expect one of %v", activity.ListTypeOptions())
					return nil
				} else {
					return err
				}
			}
			layout := "2006-01-02 15:04"
			str := arguments[3]
			date, err := time.Parse(layout, str)
			if err != nil {
				r.logger.Printf("WARN: Ensure that date format is correct. Expected format %s", layout)
				r.logger.Printf("ERROR: %s", err)
				return nil
			}

			if err := r.activityStore.Create(user, description, activityType, date); err != nil {
				return err
			}
		}
	case "q":
		r.logger.Println("WARN: Quit command takes no arguments...")
	case "h":
		r.logger.Print("WARN: Help command takes no arguments...")
	default:
		r.logger.Printf("WARN: Unknown command %s", instruction)
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
