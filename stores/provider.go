package stores

import (
	"log"
	filestore "ribbet_web/common/file_store"
	activitystore "ribbet_web/stores/activity_store"
)

func ProvideActivityStore(l *log.Logger) (activitystore.ActivityStore, error) {
	l.Println("Setting up DB and Store for Activities")

	db, err := filestore.Open(activitystore.STORE_NAME, l)
	if err != nil {
		return activitystore.ActivityStore{}, err
	}

	return activitystore.New(l, db), nil
}
