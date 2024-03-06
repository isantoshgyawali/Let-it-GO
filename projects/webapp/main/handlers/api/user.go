package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/**
Since I am using GIN, we can't visualize directly what is happening under the hood
but if I had just used bare net/http package ( Only explanation for the "func GetAllUsers()" ): This would be the way:

And due to the fact that Go 1.22 brought some good features, It's easy for me to
understand what's going on:
-- first, we are defining the userDetails of type Details at ./types.go
-- then, from net/http way we would have done

func GetAllUsers( w http.ResponseWriter, r *http.Request ) {

	//-- what parameters they are??
	// simply understanding : ResponseWriter are to write response to the client which is some form of output channel
	// and for "Request" as suggested by name: They are the request from the clients , some form of input channel

	// Now --------- Now ----------------Now

	if strings.Contains(r.Header.Get("Accept"), "application/json") {

		//-- marshalling the data to JSON
		//
		// marshalling - changing the go's data structures like map, slices into the more commonly
		// understandable formats like JSON or XML, which are easily accepted by any means on the network
		// this also refers to --> "serialization of data"

		jsonData, err := json.Marshal(userDetails) //-- use of "encoding/json" package here
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//-- Setting up the response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		//-- Finally writing JSON to the response
		w.Write(jsonData)
	}

	//serving static file

	} else if strings.Contains(c.GetHeader("Accept"), "text/html") {
		http.ServeFile(w, r, "users-data.html")
	}
}
*/

var userDetails = []*Details{
	{ID: "1", Name: "Villiers", Adress: "SA", Email: "ab17@360.com"},
	{ID: "2", Name: "McCullum", Adress: "NZ", Email: "brendon@42.com"},
}

func GetAllUsers(c *gin.Context) {
	/**
	first check if the "Accept" header contains:
	-- "application/json" then return with the JSON
	-- "text/html" then return with the html

	rememberings : anyone making request at this endpoint with the request Header that includes "application/json"
	can get userDetails so we have to consider using some middleware authorization functions/methods here but for now mehhhhh......
	*/
	if strings.Contains(c.GetHeader("Accept"), "application/json") {
		/**
		copying the userDetails, and send the
		JSON response at the endpoint
		*/
		users := append([]*Details{}, userDetails...)
		c.JSON(http.StatusOK, users)

	} else if strings.Contains(c.GetHeader("Accept"), "text/html") {

		c.HTML(http.StatusOK, "users-data.html", gin.H{
			"title":   "Users-Data-Log",
			"message": "Hello Users",
		})
	}
	// currently throws error if not matched any of these two but other format can be considered too
	// and there are many headers in http related to security, request, response..... so remeber that too
	c.AbortWithError(http.StatusNotAcceptable, fmt.Errorf("requested content type is not acceptable"))

	//Generally RestFul Api designs prioritize doing one thing at one endpoint so : you may consider that too....
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, v := range userDetails {
		if v.ID == id {
			c.JSON(http.StatusOK, v)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func CreateUser(c *gin.Context) {

}

func UpdateUserByID(c *gin.Context) {

}

func DeleteUserByID(c *gin.Context) {

}

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/", GetAllUsers)
	r.GET("/:id/", GetUserByID)

	r.POST("/", CreateUser)
	r.PUT("/:id/", UpdateUserByID)

	r.DELETE("/:id/", DeleteUserByID)
}
