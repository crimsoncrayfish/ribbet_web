package activity

import "time"

type Activity struct {
	Id           string    `json:"id"`
	Description  string    `json:"description"`
	Completed_at time.Time `json:"completion_time"`
}
