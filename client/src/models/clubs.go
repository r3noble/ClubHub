package models

type Club struct {
	Name 		string `json:"name"`
	President 	string `json:"president"`
	VP 			string `json:"VP"`
	Treasurer 	string `json:"treasurer"`
	About 		string `json:"about"`
}