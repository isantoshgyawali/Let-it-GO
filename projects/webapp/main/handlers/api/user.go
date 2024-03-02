package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


var userDetails = []details {
	{ID: "1", Name: "Villiers", Adress: "SA", Email: "ab17@360.com"},
	{ID: "2", Name: "McCullum", Adress: "NZ", Email: "brendon@42.com"},
}

func GetAllUser(c *gin.Context ) {
	//while the status is found ok , returns the map of userDetails
	// c.IndentedJSON(http.StatusOK, userDetails)
	c.HTML(http.StatusOK, "users-data.html", gin.H{
		"title" : "Users-Data-Log",
		"message" : "Hello Users",
	})

}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, v := range userDetails {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func CreateUser(c *gin.Context) {

}

func UpdateUserByID(c *gin.Context) {

}

func DeleteUserByID(c *gin.Context) {

}

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/", GetAllUser)
	r.GET("/:id/", GetUserByID)

	r.POST("/", CreateUser)
	r.PUT("/:id/", UpdateUserByID)

	r.DELETE("/:id/", DeleteUserByID)
}
