package main

import (
	"fmt"
	"log"

	"github.com/isantoshgyawali/apiWebGo/initializers"
	"github.com/isantoshgyawali/apiWebGo/handlers/router"
	"github.com/isantoshgyawali/apiWebGo/db"
)

// -- This runs even before the main function
// -- Just Loading the env for now
func init() {
	initializers.LoadEnvVar()
}

func SendQuery() {
	fmt.Printf("Let's start adding some data!")
}

func main() {
	db, err := db.ConDb()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println()
	db.Close()

	/**
	Serving the different html files at specific routes
	and then starting the server at port localohost:8081
	*/
	if err := router.RequestRouter().Run(); err != nil {
		log.Fatal(err)
	}
}