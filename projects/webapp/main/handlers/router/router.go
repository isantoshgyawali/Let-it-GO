package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	uAuth "github.com/isantoshgyawali/apiWebGo/handlers/api"
)

func handleInitialRoute(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "data-form",
	})
}

/*
*

	Well there are already multiple things to consider to make a API a RESTfulApi
	and addition to that from go - specifically gin Framework
	we got some key points to consider :

	-- use of trailing slashes should not be overlooked as I did
	   user/login/ is different than user/login

	   Yes, that is what it is : Interms of Gin I guess
	   Framework redirect the user/login to user/login/ itself though

	   but considering the consistent approach of using the "/" while using Gin
	   is I guess a better choice ( "I am a very newbie so don't take me seriously" )

	-- Don't create the multiple Routers here and there without understanding
	   what you are trying to do : Atleast I learnt from here what mess I made earlier

	-- Grouping Routes of similar type makes code more fluid to grasp

	-- You can create a engine of two types
		a. gin.Default() : This comes with two middlewares preconfigured logger && recovery
		b. gin.New() : This just comes with no middlewares , completely freshed

		while gin.New() is abit faster than gin.Default in general
		about 50-500 microseconds here and there

		but anyway gin.Default() is better to go with as those two middlewares helps in debugging
*/
func RequestRouter() *gin.Engine {
	r := gin.Default()

	serveFiles(r)

	r.GET("/", handleInitialRoute)
	userGroup := r.Group("/user/")
	uAuth.UserRoutes(userGroup)

	orgGrop := r.Group("/org/")
	uAuth.OrgRoutes(orgGrop)

	return r
}

func serveFiles(r *gin.Engine){
	// r.StaticFS("/form", gin.Dir("../ui", true))
	/**
	 loading specific files - not good approach 
	 r.LoadHTMLFiles("../ui/users-data.html") 

	 rather , serving the filePatterns from fs - DRY
	*/
	r.LoadHTMLGlob("../ui/*")
}
