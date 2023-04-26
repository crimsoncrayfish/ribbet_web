package activitystore

import (
	"log"
	"ribbet_web/models/activity"
	"time"
)

type ActivityStore struct {
	l *log.Logger
}

func New(l *log.Logger) ActivityStore {
	return ActivityStore{
		l: l,
	}
}

func (as ActivityStore) Create(description string, completed_at time.Time) error {
	return nil
}

func (as ActivityStore) List() ([]activity.Activity, error) {
	return nil, nil
}
