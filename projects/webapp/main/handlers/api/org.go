package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var orgDetails = []*Details { 
	{ID: "1", Name: "TESLA", Adress: "USA", Email: "tesla@test.com"},
	{ID: "2", Name: "TWITTER", Adress: "USA", Email: "tesla@test.com"},
	{ID: "3", Name: "BORING", Adress: "USA", Email: "boring@test.com"},
}

func GetAllOrg(c *gin.Context) {
	if strings.Contains(c.GetHeader("Accept"), "application/json") {
		users := append([]*Details{}, orgDetails...)
		c.JSON(http.StatusOK, users)

	} else if strings.Contains(c.GetHeader("Accept"), "text/html") {

		c.HTML(http.StatusOK, "org-data.html", gin.H{
			"title":   "org | Log",
			"message": "org-Data-Log",
		})
	}

	c.AbortWithError(http.StatusNotAcceptable, fmt.Errorf("requested content type is not acceptable"))
}

func GetOrgByID(c *gin.Context) {
	id := c.Param("id")

	for _, v := range orgDetails{
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "org not found"})
}

func CreateOrg(c *gin.Context) {

}

func UpdateOrgByID(c *gin.Context) {

}

func DeleteOrgByID(c *gin.Context) {

}

func OrgRoutes(r *gin.RouterGroup) {
	r.GET("/", GetAllOrg)
	r.GET("/:id/", GetOrgByID)

	r.POST("/", CreateOrg)
	r.PUT("/:id/", UpdateOrgByID)

	r.DELETE("/:id/", DeleteOrgByID)
}
