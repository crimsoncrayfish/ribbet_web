package filestore

import (
	"errors"
	"io"
	"log"
	"os"
)

func Open(storeName string, l *log.Logger) (*os.File, error) {
	filePath := "./tmp/tmp." + storeName + ".json"

	l.Printf("Reading/creating file at \"%s\"\n", filePath)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	l.Println("Validating file not empty")
	if err := isFileValidJson(file); err != nil {
		return nil, err
	}

	return file, nil
}

const EMPTY_STORE_VALUE = "[]"

func isFileValidJson(file *os.File) error {
	file.Seek(0, 0)

	_, err := file.Read(make([]byte, 1))
	if err != nil {
		if errors.Is(err, io.EOF) {
			//handle empty file
			file.Write([]byte(EMPTY_STORE_VALUE))
		} else {
			return err
		}
	}
	file.Seek(0, 0)
	return nil
}
