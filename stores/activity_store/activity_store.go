package activitystore

import (
	"encoding/json"
	"log"
	"os"
	"ribbet_web/models/activity"
	"time"

	"github.com/google/uuid"
)

const STORE_NAME = "activities"

type ActivityStore struct {
	l      *log.Logger
	dbFile *os.File
}

func New(l *log.Logger, db *os.File) ActivityStore {
	return ActivityStore{
		l:      l,
		dbFile: db,
	}
}

func (as ActivityStore) Delete(id string) error {
	activities, err := as.readAll()
	if err != nil {
		return err
	}

	index, err := indexOfId(activities, id)
	if err != nil {
		return err
	}
	activities[index] = activities[len(activities)-1]
	activities = activities[:len(activities)-1]

	if err := as.writeAll(activities); err != nil {
		return err
	}

	return nil
}

func (as ActivityStore) Create(description string, completed_at time.Time) error {
	activity := activity.Activity{
		Description:  description,
		Completed_at: completed_at,
		Id:           uuid.New().String(),
	}

	activities, err := as.readAll()
	if err != nil {
		return err
	}

	activities = append(activities, activity)

	if err := as.writeAll(activities); err != nil {
		return err
	}

	return nil
}

func (as ActivityStore) List() ([]activity.Activity, error) {
	activities, err := as.readAll()
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (as ActivityStore) readAll() ([]activity.Activity, error) {
	//set file cursor to start of file
	_, err := as.dbFile.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	//read activities from file
	var activities []activity.Activity
	err = json.NewDecoder(as.dbFile).Decode(&activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (as ActivityStore) writeAll(activities []activity.Activity) error {
	//set file cursor to start of file
	_, err := as.dbFile.Seek(0, 0)
	if err != nil {
		return err
	}

	//read activities from file
	err = json.NewEncoder(as.dbFile).Encode(activities)
	if err != nil {
		return err
	}

	return nil
}

func indexOfId(activities []activity.Activity, id string) (int, error) {
	for i, activity := range activities {
		if activity.Id == id {
			return i, nil
		}
	}
	return 0, NotFoundError
}
