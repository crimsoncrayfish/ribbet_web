package typeconvert

import (
	"fmt"
	"time"
)

func StringToDate(input string) (time.Time, error) {
	layout := "2006-01-02 15:04"
	date, err := time.Parse(layout, input)
	if err != nil {
		layout2 := "2006-01-02"
		date, err2 := time.Parse(layout2, input)
		if err2 != nil {
			fmt.Printf("WARNING: Ensure that date format is correct. Expected format %s or %s\n", layout, layout2)
			return time.Now(), ErrStringToDateConvertFailed
		}
		return date, nil
	}
	return date, nil
}
