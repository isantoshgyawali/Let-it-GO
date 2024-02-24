package main

import (
	"fmt"
	"time"

	// "github.com/gin-gonic/gin"
	"github.com/isantoshgyawali/apiWebGo/handlers/router"
)

func main() {
	s := time.Now()
	fmt.Println("hello there")

	// r := gin.Default()
	r := router.RequestRouter()

    fmt.Println(r)

	e := time.Since(s)
	fmt.Printf("\n --------------------------------\n#  Program executed in %v  #\n --------------------------------\n\n", e)

	// r.Run(":8080")

}
