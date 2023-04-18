package activity

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type ActivityFileStore struct {
	database *os.File
	l        *log.Logger
}

func NewActivityFileStore(newDatabase *os.File, l *log.Logger) ActivityFileStore {
	return ActivityFileStore{newDatabase, l}
}

func (afs ActivityFileStore) Create(user, description string, activityType Type, completed time.Time) error {
	afs.l.Println("Creating new Activity")
	activities, err := afs.readAll()
	if err != nil {
		return err
	}
	newId := uuid.NewString()
	activities = append(activities, Activity{
		Id:           newId,
		UserName:     user,
		Description:  description,
		ActivityType: activityType,
		CompletedAt:  completed,
	})
	if err := afs.writeAll(activities); err != nil {
		return err
	}
	afs.l.Printf("Created new Activity with id '%s'", newId)
	return nil
}

func (afs ActivityFileStore) Update(id, user, description string, activityType Type, completed time.Time) error {
	afs.l.Printf("Updating Activity with ID '%s'", id)
	activities, err := afs.readAll()
	if err != nil {
		return err
	}
	index, err := getIndexById(activities, id)
	if err != nil {
		return err
	}

	updatedActivity := Activity{
		Id:           id,
		UserName:     user,
		Description:  description,
		ActivityType: activityType,
		CompletedAt:  completed,
	}
	activities = append(append(activities[:index], updatedActivity), activities[index+1:]...)

	if err := afs.writeAll(activities); err != nil {
		return err
	}
	afs.l.Printf("Updated new Activity with id '%s'", updatedActivity.Id)
	return nil
}

func (afs ActivityFileStore) Delete(id string) error {
	afs.l.Printf("Deleting Activity with ID '%s'", id)
	activities, err := afs.readAll()
	if err != nil {
		return err
	}
	index, err := getIndexById(activities, id)
	if err != nil {
		return err
	}

	activities = append(activities[:index], activities[index+1:]...)

	if err := afs.writeAll(activities); err != nil {
		return err
	}
	afs.l.Printf("Deleted Activity with ID '%s'", id)
	return nil
}

func (afs ActivityFileStore) List() ([]Activity, error) {
	afs.l.Println("List Activities")

	activities, err := afs.readAll()
	if err != nil {
		return nil, err
	}
	afs.l.Println("Listed Activities")
	return activities, nil
}

func (afs ActivityFileStore) readAll() ([]Activity, error) {
	//set cursor to first position
	afs.database.Seek(0, 0)

	var activities []Activity

	//read and decode content of database into activities
	err := json.NewDecoder(afs.database).Decode(&activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (afs ActivityFileStore) writeAll(activities []Activity) error {
	if err := afs.database.Truncate(0); err != nil {
		return err
	}
	if _, err := afs.database.Seek(0, 0); err != nil {
		return err
	}
	if err := json.NewEncoder(afs.database).Encode(activities); err != nil {
		return err
	}

	return nil
}

func getIndexById(activities []Activity, id string) (int, error) {
	index := -1
	for i := 0; i < len(activities); i++ {
		activity := activities[i]
		if activity.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return index, ErrIdNotFound
	}

	return index, nil
}
