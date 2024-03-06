package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var orgDetails = []*Details { 
	{ID: "1", Name: "TESLA", Adress: "USA", Email: "tesla@test.com"},
	{ID: "2", Name: "TWITTER", Adress: "USA", Email: "tesla@test.com"},
	{ID: "1", Name: "BORING", Adress: "USA", Email: "boring@test.com"},
}

func GetAllOrg(c *gin.Context) {
	c.HTML(http.StatusOK, "org-data.html", gin.H{
		"title": "Org-Data-log",
		"message": "Hello Org",
	})

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
