package api

import (
	//"fmt"
	//"net/http"

	"github.com/gin-gonic/gin"
)

// type org struct {
//  ID     string `json:"id"`
//  Name   string `json:"name"`
//  Adress string `json:"address"`
//  Email  string `json:"email"`
//}
//

func GetAllOrg(c *gin.Context) {

}

func GetOrgByID(c *gin.Context) {

}

func CreateOrg(c *gin.Context) {

}

func UpdateOrgByID(c *gin.Context) {

}

func DeleteOrgByID(c *gin.Context) {

}

func OrgRoutes() *gin.RouterGroup {
	ur := gin.New()

	orgGroup:= ur.Group("/org")
	{
		orgGroup.GET("/", GetAllOrg)
		orgGroup.GET("/:id", GetOrgByID)

		orgGroup.POST("/", CreateOrg)
		orgGroup.PUT("/:id", UpdateOrgByID)

		orgGroup.DELETE("/:id", DeleteOrgByID)
	}

	return orgGroup
}
