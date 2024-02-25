package main

import (
	"github.com/isantoshgyawali/apiWebGo/handlers/router"
)

func main() {
	//starting the router
	router.RequestRouter().Run(":8081")
}
