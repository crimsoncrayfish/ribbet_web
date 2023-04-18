package activity

import (
	"strings"
	"time"
)

type Activity struct {
	Id           string    `json:"id"`
	UserName     string    `json:"user_name"`
	Description  string    `json:"description"`
	CompletedAt  time.Time `json:"completed_at"`
	ActivityType Type      `json:"activity_type"`
}

type Type int

const (
	None Type = iota
	Run
	Walk
	Music
	Art
	Course
)

func ListTypeOptions() [5]string {
	return [5]string{Run.ToString(), Walk.ToString(), Music.ToString(), Art.ToString(), Course.ToString()}
}

func (t Type) ToString() string {
	switch t {
	case Run:
		return "Run"
	case Walk:
		return "Walk"
	case Music:
		return "Music"
	case Art:
		return "Art"
	case Course:
		return "Course"
	}
	return ""
}

func ToType(input string) (Type, error) {
	switch strings.ToLower(input) {
	case "run":
		return Run, nil
	case "walk":
		return Walk, nil
	case "music":
		return Music, nil
	case "art":
		return Art, nil
	case "course":
		return Course, nil
	default:
		return None, ErrActivityTypeNotFound
	}
}

func (value Activity) IsMatch(expected Activity) bool {
	if (value.Id != expected.Id) ||
		(value.Description != expected.Description) ||
		(value.UserName != expected.UserName) ||
		(!value.CompletedAt.Equal(expected.CompletedAt)) ||
		(value.ActivityType != expected.ActivityType) {
		return false
	}
	return true
}
