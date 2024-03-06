package router
import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	uAuth "github.com/isantoshgyawali/apiWebGo/handlers/api"
)

//-- If the client accepts HTML, serve the HTML file
//-- Otherwise, serve the JSON data
func handleInitialRoute(c *gin.Context) {
	if strings.Contains(c.GetHeader("Accept"), "text/html") {
		c.HTML(http.StatusOK, "index.html", nil)
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"title" : "Home | Add data",
			"message": "what's up?",
		})
	}
}

/**
	Well there are already multiple things to consider to make a API a RESTfulApi
	and addition to that from go - specifically gin Framework
	we got some key points to consider :

	-- use of trailing slashes should not be overlooked as I did
	   user/login/ is different than user/login

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

	r.GET("/", handleInitialRoute) //-- index page
	userGroup := r.Group("/user/") //-- users | groups/page
	uAuth.UserRoutes(userGroup)	

	orgGrop := r.Group("/org/")    //-- org | groups/page
	uAuth.OrgRoutes(orgGrop)

	return r
}

func serveFiles(r *gin.Engine) {

	/**
		first serving the other static files like:
		js, css, image , or other media using StaticFS from gin

		then: UPDATE the link url in your html file to import them
		in this case /app/static/style.css in ../ui/static/index.html
	*/
	r.StaticFS("/app/scripts/", gin.Dir("../ui/scripts/", true))
	r.StaticFS("/app/static/", gin.Dir("../ui/static/", true))

	/**
		loading specific files - not good approach
		r.LoadHTMLFiles("../ui/static/users-data.html")

		rather , serving the filePatterns from fs - DRY
	*/
	r.LoadHTMLGlob("../ui/static/*")
}
