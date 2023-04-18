package activity_test

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"ribbit_web/activity"
	"ribbit_web/core/logx"
)

func TestActivityFileStore(t *testing.T) {
	initialData := []activity.Activity{
		{
			Id:           "one",
			UserName:     "Blue",
			Description:  "Do the thing",
			CompletedAt:  time.Now(),
			ActivityType: activity.Art,
		},
		{
			Id:           "two",
			UserName:     "Red",
			Description:  "Do the opposite",
			CompletedAt:  time.Now(),
			ActivityType: activity.Course,
		},
	}
	initialListLength := len(initialData)

	initialDataJson, err := json.Marshal(initialData)
	if err != nil {
		t.Fatalf("Failed to set up initial data json with error %v", err)
	}

	tempDB, cleanup := createTempDb(t, string(initialDataJson))
	defer cleanup()

	l := logx.New("Tests", true)
	activityStore := activity.NewActivityFileStore(tempDB, l)

	t.Run("Create and list activities", func(t *testing.T) {
		err := activityStore.Create(
			"Someone",
			"Testing Create and List",
			activity.Music,
			time.Now(),
		)
		assertNilErr(t, err)

		activities, err := activityStore.List()
		assertNilErr(t, err)
		assertListLength(t, activities, initialListLength+1)
	})

	t.Run("Delete and list activities", func(t *testing.T) {
		//get initial list
		activities, err := activityStore.List()
		assertNilErr(t, err)
		newListLength := len(activities)

		//delete first item in DB by its id
		idToBeDeleted := activities[0].Id
		err = activityStore.Delete(idToBeDeleted)
		assertNilErr(t, err)

		//get new list
		activities, err = activityStore.List()
		assertNilErr(t, err)
		assertListLength(t, activities, newListLength-1)
		assertListDoesntContainId(t, activities, idToBeDeleted)
	})

	t.Run("Delete non existant id should fail", func(t *testing.T) {
		//delete first item in DB by its id
		idToBeDeleted := "Id that doesnt exist"
		err := activityStore.Delete(idToBeDeleted)
		if !activity.IsNotFoundError(err) {
			t.Errorf("Expected 'IsNotFoundError' but got '%s'", err)
		}
	})

	t.Run("Update and list activities", func(t *testing.T) {
		//get initial list
		activities, err := activityStore.List()
		assertNilErr(t, err)
		newListLength := len(activities)

		//update first activity
		first := activities[0]
		updatedActivity := activity.Activity{
			Id:           first.Id,
			UserName:     "Updated",
			Description:  "Updated",
			CompletedAt:  time.Now(),
			ActivityType: activity.Music,
		}
		err = activityStore.Update(
			updatedActivity.Id,
			updatedActivity.UserName,
			updatedActivity.Description,
			updatedActivity.ActivityType,
			updatedActivity.CompletedAt)
		assertNilErr(t, err)

		activities, err = activityStore.List()
		assertNilErr(t, err)
		//list size shouldnt increase
		assertListLength(t, activities, newListLength)
		assertActivityExists(t, updatedActivity, activities)
	})
}

func assertNilErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected nil err, but got '%s'", err)
	}
}

func assertListLength(t testing.TB, activities []activity.Activity, expectedLength int) {
	t.Helper()
	if len(activities) != expectedLength {
		t.Errorf("Expected '%d' item/s but got '%d'", expectedLength, len(activities))
	}
}

func assertListDoesntContainId(t testing.TB, activities []activity.Activity, missingId string) {
	t.Helper()
	for i := 0; i < len(activities)-1; i++ {
		activity := activities[i]
		if activity.Id == missingId {
			t.Errorf("Expected id '%s' to be missing from activities, but it was present", missingId)
		}
	}
}

func assertActivityExists(t testing.TB, expectedActivity activity.Activity, activities []activity.Activity) {
	t.Helper()

	found := false
	for i := 0; i < len(activities)-1; i++ {
		isMatch := expectedActivity.IsMatch(activities[i])
		if isMatch {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected Activity was not found in list")
		t.Error("-------------Expected Activity-----------------")
		t.Error(expectedActivity)
		t.Error("-------------Full List-------------------------")
		t.Error(activities)
		t.Error("-----------------------------------------------")
	}
}

func createTempDb(t testing.TB, inputData string) (*os.File, func()) {
	t.Helper()

	tempDB, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("Could not create temp DB file for tests, %v", err)
	}

	if _, err := tempDB.Write([]byte(inputData)); err != nil {
		t.Fatalf("Failed to write data to temp DB with error %v", err)
	}

	removeFile := func() {
		tempDB.Close()
		os.Remove(tempDB.Name())
	}

	return tempDB, removeFile
}
