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

var orgDetails = []*Details{}

func LoadOrgsFromDB() ([]*Details, error) {
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

	stmt, err := db.Prepare("SELECT id, name, address, email FROM root.orgs;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orgs []*Details
	for rows.Next() {
		org := &Details{}
		if err := rows.Scan(&org.ID, &org.Name, &org.Adress, &org.Email); err != nil {
			log.Fatal(err)
		}
		orgs = append(orgs, org)
	}

	return orgs, nil
}

func GetAllOrg(c *gin.Context) {
	orgs, err := LoadOrgsFromDB()
	if err != nil {
		fmt.Printf("There was a error Loading the data from the db\n")
		log.Fatal(err)
	}

	if strings.Contains(c.GetHeader("Accept"), "application/json") {
		c.JSON(http.StatusOK, orgs)

	} else if strings.Contains(c.GetHeader("Accept"), "text/html") {

		c.HTML(http.StatusOK, "org-data.html", gin.H{
			"title":   "orgs | Log",
			"message": "orgs-Data-Log",
		})
	}
}

func GetOrgByID(c *gin.Context) {
	id := c.Param("id")

	orgs, err := LoadOrgsFromDB()
	if err != nil {
		fmt.Printf("There was error loading the data from db")
		log.Fatal(err)
	}

	for _, org := range orgs {
		if org != nil && strconv.Itoa(org.ID) == id {
			c.JSON(http.StatusOK, org)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "org not found"})
}

func CreateOrg(c *gin.Context) {
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

	query, err := os.Open("db/pq/org/CreateOrg.sql")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer query.Close()

	cnt, err := io.ReadAll(query)
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare(string(cnt))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(formData.Name, formData.Adress, formData.Email)
	if err != nil {
		log.Printf("Error inserting org: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create org"})
		return
	}

	// adding the formData to the orgDetails slice
	orgDetails = append(orgDetails, formData)

	c.JSON(http.StatusOK, gin.H{"message": "org Created Successfully"})
}

func DeleteOrgByID(c *gin.Context) {
	id := c.Param("id")

	db, err := initializers.ConDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	stmt, err := db.Prepare("DELETE FROM root.orgs WHERE id = $1")
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
		c.JSON(http.StatusNotFound, gin.H{"message": "org not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "org deleted"})

}

func DeleteAllOrg(c *gin.Context){
	db, err := initializers.ConDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	stmt, err := db.Prepare("DELETE from root.orgs;")
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
		c.JSON(http.StatusNotFound, gin.H{"message": "orgs not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "orgs deleted"})
}

func OrgRoutes(r *gin.RouterGroup) {
	r.GET("/", GetAllOrg)
	r.GET("/:id/", GetOrgByID)

	r.POST("/", CreateOrg)
	r.DELETE("/:id/", DeleteOrgByID)
	r.DELETE("/", DeleteAllOrg)

	// r.PUT("/:id/", UpdateOrgByID) //-- nothing to do here for now
}
