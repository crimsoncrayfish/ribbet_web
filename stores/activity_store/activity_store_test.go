package activitystore_test

import (
	logr "ribbet_web/common/log_r"
	activitystore "ribbet_web/stores/activity_store"
	"testing"
	"time"
)

func TestActivityStore(t *testing.T) {
	l := logr.New("Testing")
	store := activitystore.New(l)
	t.Run("Create new Activity", func(t *testing.T) {
		activities, err := store.List()
		if err != nil {
			t.Error(err)
		}
		initialCount := len(activities)
		store.Create("What up", time.Now())

		activities, err = store.List()
		if err != nil {
			t.Error(err)
		}

		if len(activities)-initialCount != 1 { //1 new item created
			t.Errorf("expected new list to be 1 longer than old list; New list length: %d, Old list length: %d", len(activities), initialCount)
		}
	})
}
