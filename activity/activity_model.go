package activity

import "time"

type Activity struct {
	UID         string    `json:"id"`
	Description string    `json:"description"`
	CompletedAt time.Time `json:"completed_at"`
	CompletedBy string    `json:"completed_by"`
}
