package activitystore_test

import (
	"encoding/json"
	"os"
	logr "ribbet_web/common/log_r"
	"ribbet_web/models/activity"
	activitystore "ribbet_web/stores/activity_store"
	"testing"
	"time"
)

func TestActivityStore(t *testing.T) {
	l := logr.New("Testing")

	initialActivities := []activity.Activity{
		{
			Id:           "One",
			Description:  "Initial data 1",
			Completed_at: time.Now(),
		},
		{
			Id:           "Two",
			Description:  "Initial data 2",
			Completed_at: time.Now(),
		},
	}

	initialData, err := json.Marshal(initialActivities)
	if err != nil {
		t.Fatalf("Failed to set up initial data for tests with error %s", err)
	}

	db, cleanupFunc := createTempFile(t, string(initialData))
	defer cleanupFunc()

	store := activitystore.New(l, db)

	t.Run("Create new Activity", func(t *testing.T) {
		activities, err := store.List()
		if err != nil {
			t.Fatal(err)
		}
		initialCount := len(activities)
		if err := store.Create("What up", time.Now()); err != nil {
			t.Fatal(err)
		}

		activities, err = store.List()
		if err != nil {
			t.Fatal(err)
		}

		if len(activities)-initialCount != 1 { //1 new item created
			t.Errorf("expected new list to be 1 longer than old list; New list length: %d, Old list length: %d", len(activities), initialCount)
		}
	})
	t.Run("Delete Activity", func(t *testing.T) {
		activities, err := store.List()
		if err != nil {
			t.Fatal(err)
		}
		initialCount := len(activities)
		idToDelete := activities[0].Id

		if err := store.Delete(idToDelete); err != nil {
			t.Fatal(err)
		}

		activities, err = store.List()
		if err != nil {
			t.Fatal(err)
		}

		if len(activities)-initialCount != -1 { //1 item deleted
			t.Errorf("expected new list to be 1 shorter than old list; New list length: %d, Old list length: %d", len(activities), initialCount)
		}

	})
	t.Run("Delete Activity not exists", func(t *testing.T) {
		activities, err := store.List()
		if err != nil {
			t.Fatal(err)
		}
		initialCount := len(activities)
		idToDelete := activities[0].Id

		if err := store.Delete(idToDelete); err != nil {
			t.Fatal(err)
		}

		activities, err = store.List()
		if err != nil {
			t.Fatal(err)
		}

		if len(activities)-initialCount != -1 { //1 item deleted
			t.Errorf("expected new list to be 1 shorter than old list; New list length: %d, Old list length: %d", len(activities), initialCount)
		}

	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	file, err := os.CreateTemp("", "temp.db")
	if err != nil {
		t.Fatalf("Failed to create temp file with err %s", err)
	}
	if _, err := file.Write([]byte(initialData)); err != nil {
		t.Fatalf("Failed to initialise starting data with err %s", err)
	}

	cleanupFunc := func() {
		file.Close()
		os.Remove(file.Name())
	}

	return file, cleanupFunc
}
