package activity

import "time"

type Activity struct {
	id           string
	description  string
	completed_at time.Time
}
