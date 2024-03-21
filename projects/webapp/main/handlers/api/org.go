package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var orgDetails = []*Details { 
	{ Name: "TESLA", Adress: "USA", Email: "tesla@test.com"},
	{ Name: "TWITTER", Adress: "USA", Email: "tesla@test.com"},
	{ Name: "BORING", Adress: "USA", Email: "boring@test.com"},
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
	for _, org := range orgDetails{
		if strconv.Itoa(org.ID) == id {
			c.IndentedJSON(http.StatusOK, org)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "org not found"})
}

func CreateOrg(c *gin.Context) {
	formData, err := GetFormData(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error occured": err.Error()})
	}
	orgDetails = append(orgDetails, formData)
	fmt.Println(*orgDetails[len(orgDetails)-1])
}

func DeleteOrgByID(c *gin.Context) {
	id := c.Param("id")
	for _, org := range orgDetails{
		if strconv.Itoa(org.ID) == id {
			c.IndentedJSON(http.StatusOK, org)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "org not found"})
}

func OrgRoutes(r *gin.RouterGroup) {
	r.GET("/", GetAllOrg)
	r.GET("/:id/", GetOrgByID)

	r.POST("/", CreateOrg)
	r.DELETE("/:id/", DeleteOrgByID)

	// r.PUT("/:id/", UpdateOrgByID) //-- nothing to do here for now
}
