package models

type Event struct {
	Club string `json:"club"`
	Event string `json:"event"`
	Date string `json:"date"`
	startTime string `json:"startTime"`
	endTime string `json:"endTime"`
	Users string `json:"users"`
}