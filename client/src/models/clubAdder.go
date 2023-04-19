package models

type ClubAdder struct {
	//string of ID correlating to user
	ID string `json:"id"`
	//name of club to add
	Name string `json:"name"`
}