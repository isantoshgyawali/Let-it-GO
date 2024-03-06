package main
import (
	"log"

	"github.com/isantoshgyawali/apiWebGo/handlers/router"
	"github.com/isantoshgyawali/apiWebGo/initializers"
)

//-- This runs even before the main function 
//-- Just Loading the env for now
func init(){
	initializers.LoadEnvVar()
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
