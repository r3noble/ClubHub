package models

type Event struct {
	Club string `json:"club"`
	Event string `json:"event"`
	Date string `json:"date"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	Users string `json:"users"`
}