package filestore

import (
	"io"
	"log"
	"os"
)

func InitDatabaseFile(l *log.Logger, locationEnv string) (*os.File, error) {
	//get location
	location := os.Getenv(locationEnv)
	if location == "" {
		location = "tmp.store.json"
	}

	l.Printf("File Store location is '%s", location)

	//open file
	l.Printf("Opening File Store")
	store, err := os.OpenFile(location, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	//json.Marchel doesnt work w/ empty json file
	// ensure file store is not empty
	if _, err := store.Read(make([]byte, 1)); err != nil {
		if err == io.EOF {
			l.Printf("Bootstrapping new File Store")
			//write empty json to brand new File Store
			store.Write([]byte("[]"))

			//reset cursor to start of file
			store.Seek(0, 0)
		}
	}

	return store, nil
}
