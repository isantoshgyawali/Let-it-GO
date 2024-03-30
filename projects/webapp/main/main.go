package main

import (
	"fmt"
	"log"

	"github.com/isantoshgyawali/apiWebGo/handlers/router"
	"github.com/isantoshgyawali/apiWebGo/initializers"
)

// -- This runs even before the main function
//-- loading env
//-- connecting to the db
//-- creating tables if not already present

// ( THERE WILL BE NEED OF LINUX MACHINE TO RUN THIS CODE, I ASSUME
// 	 THERE WILL BE MESS TO SETUP THE Postgres ON WINDOWS FROM ROCK BOTTOM , WELL I DON'T KNOW TBH )

//But, creating the db using pgAdmin or ...
//THEN, connecting directly without any setup will be the choice I guess

//DON'T KNOW ABOUT WINDOWS AND PgAdmin Stuff
//JUST SIMPLE RECOMMENDATION TO ANYONE WHO IS READING THIS "LEARN LINUX"
//PRODUCTION SERVER WON'T BE RUNNING ON WINDOWS, MOST OF THE TIME if not ALL
func init() {
	initializers.LoadEnvVar()
	initializers.ConDb()

	//-- akshually :) we don't need this two initializers as
	//-- On production enviroment - we will be creating db or will have already created
	//-- entire tables, schema structure seperately and if not there would be migrations scripts
	//
	//-- BUT for the development enviromet - it could be a good practice as it maintains consistency while development
	if err := initializers.CreateOrgTable(); err != nil {
		fmt.Println(err)
	}

	if err := initializers.CreateUserTable(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	/**
	Serving the different html files at specific routes
	and then starting the server at port localohost:8081
	*/
	if err := router.RequestRouter().Run(); err != nil {
		log.Fatal(err)
	}
}
