package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/isantoshgyawali/apiWebGo/initializers"
)

/*
*
Since I am using GIN, we can't visualize directly what is happening under the hood
but if I had just used bare net/http package ( Only explanation for the "func GetAllUsers()" ): This would be the way:

And due to the fact that Go 1.22 brought some good features, It's easy for me to
understand what's going on:
-- first, we are defining the userDetails of type Details at ./types.go
-- then, from net/http way we would have done

func GetAllUsers( w http.ResponseWriter, r *http.Request ) {

		//-- what are those parameters??
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

		//-- serving static file
		} else if strings.Contains(c.GetHeader("Accept"), "text/html") {
			http.ServeFile(w, r, "users-data.html")
		}
	}

	// at ../router/router.go we should define the endpoints and start the server :
	// ( we don't see any logs like in gin which uses middleware named logger(), so just print starting.... )

	// Defining Port and checking if there is env storing them:

	port := os.Getenv("PORT") //-- using os package from go std.library
	if port == "" {
		port = "8080" //-- default port if port are not defined/available in env
	}

	http.Handlefunc("/user/", GetAllUser
	fmt.Println("starting.....")
	http.ListenAndServe(":"+port, handler)
*/
var userDetails = []*Details{}

func LoadUsersFromDB() ([]*Details, error) {
	// getting the data from the database
	/**
	-- Connecting to the db
	-- Preaparing the query
	-- Executing the query
	*/
	db, err := initializers.ConDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, name, address, email FROM root.users;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*Details
	for rows.Next() {
		user := &Details{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Adress, &user.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, nil
}

func GetAllUsers(c *gin.Context) {
	users, err := LoadUsersFromDB()
	if err != nil {
		fmt.Printf("There was a error Loading the data from the db\n")
		log.Fatal(err)
	}
	/**
	first check if the "Accept" header contains:
	-- "application/json" then return with the JSON
	-- "text/html" then return with the html

	rememberings : anyone making request at this endpoint with the request Header that includes "application/json"
	can get userDetails so we have to consider using some middleware authorization functions/methods here but for now mehhhhh......
	*/
	if strings.Contains(c.GetHeader("Accept"), "application/json") {
		//-- send the response at the endpoint
		c.JSON(http.StatusOK, users)

	} else if strings.Contains(c.GetHeader("Accept"), "text/html") {

		c.HTML(http.StatusOK, "users-data.html", gin.H{
			"title":   "Users | Log",
			"message": "Users-Data-Log",
		})
	}

	// we can throw error if not matched any of these but other format can be considered too
	// and there are many headers in http related to security, request, response.....
	//
	//-------------------------    throw error if necessary   -------------------------------------
	// c.AbortWithError(http.StatusNotAcceptable, fmt.Errorf("requested content type is not acceptable"))
	//
	// self-note : Generally RestFul Api designs prioritize doing one thing at one endpoint so : you might consider that too....
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	users, err := LoadUsersFromDB()
	if err != nil {
		fmt.Printf("There was error loading the data from db")
		log.Fatal(err)
	}

	for _, user := range users {
		if user != nil && strconv.Itoa(user.ID) == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func CreateUser(c *gin.Context) {
	formData, err := GetFormData(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error occured": err.Error()})
	}

	// adding the formData to the database
	db, err := initializers.ConDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// "os.Open()" method returns a pointer to the file Object ie. *FILE
	// this don't allows directly to read the file but the file can be read using
	// other methods such as "io.ReadAll()" which returns the file contents in the byte slice format
	// which then could be converted to a understandable string using the "string()" method

	// for the simple queries like this and which takes the user input
	// it's a recommended approach to write the queries in the backend code itself which
	// helps in avoiding the sql injection properly as it gives more control rather reading the file
	query, err := os.Open("db/pq/user/CreateUser.sql")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer query.Close()

	cnt, err := io.ReadAll(query)
	if err != nil {
		log.Fatal(err)
	}
	// cnt reads the file in the byte slice form
	// use the string method to modify the data into understandable strings
	stmt, err := db.Prepare(string(cnt))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(formData.Name, formData.Adress, formData.Email)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// adding the formData to the userDetails slice
	userDetails = append(userDetails, formData)

	c.JSON(http.StatusOK, gin.H{"message": "User Created Successfully"})
}
 
func DeleteAllUser(c *gin.Context){
	db, err := initializers.ConDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	stmt, err := db.Prepare("DELETE from root.users;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
	}
	defer stmt.Close()

	result, err := stmt.Exec()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowChanged, err := result.RowsAffected()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if rowChanged == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")

	db, err := initializers.ConDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	stmt, err := db.Prepare("DELETE FROM root.users WHERE id = $1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowChanged, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowChanged == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/", GetAllUsers)
	r.GET("/:id/", GetUserByID)
	r.POST("/", CreateUser)
	// r.PUT("/:id/", UpdateUserByID) //-- nothing to update currently
	r.DELETE("/:id/", DeleteUserByID)
	r.DELETE("/", DeleteAllUser)
}
