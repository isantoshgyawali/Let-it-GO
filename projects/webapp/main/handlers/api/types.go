package api

//-- omitempty is added to neglect the ID while Binding JSON
//-- as it is not available on the context itself 
type Details struct {
	ID     		int 	`json:"id,omitempty"` 
	Name   		string	`json:"name"`		  
	Adress 		string	`json:"address"`
	Email		string	`json:"email"`
	Type		string	`json:"type"`
}