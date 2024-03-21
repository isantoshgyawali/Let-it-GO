package api

type Details struct {
	ID     		int 	`json:"id,omitempty"`
	Name   		string	`json:"name"`
	Adress 		string	`json:"address"`
	Email		string	`json:"email"`
	Type		string	`json:"type"`
}