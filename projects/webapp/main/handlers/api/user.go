package api

import (
	//"fmt"
	//"net/http"
	//"fmt"

	"github.com/gin-gonic/gin"
)

// type user struct {
//  ID     string `json:"id"`
//  Name   string `json:"name"`
//  Adress string `json:"address"`
//  Email  string `json:"email"`
//}
//

func GetAllUser(c *gin.Context) {

}

func GetUserByID(c *gin.Context) {

}

func CreateUser(c *gin.Context) {

}

func UpdateUserByID(c *gin.Context) {

}

func DeleteUserByID(c *gin.Context) {

}

func UserRoutes() *gin.RouterGroup {
	// creates blanck engine with no preConfigured middleware
	// also, gin.Default() can be used : it comes with two preConfigured middlewares
	// loggers && Recovery
	ur := gin.New()

	userGroup := ur.Group("/user")
	{
		userGroup.GET("/",GetAllUser)
		userGroup.GET("/:id",GetUserByID)

		userGroup.POST("/",CreateUser)
		userGroup.PUT("/:id",UpdateUserByID)

		userGroup.DELETE("/:id",DeleteUserByID)
	}

	return userGroup
}
