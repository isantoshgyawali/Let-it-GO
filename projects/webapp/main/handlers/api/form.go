package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetFormData(c *gin.Context) (*Details, error) {
	//-- Parseform is used when traditional form data is sent over the http request using html directly
	//-- but as we are sending json using axios we will be using BindJSON
	//
	// if err := c.Request.ParseForm(); err != nil {
	// 	fmt.Println("Error Parsing the form data", err)
	// 	return nil, err
	// }
	// fmt.Println("Form Data: ", c.Request.PostForm)

	formData:= &Details{}
	/**
	explicitly allocating so that it doesn't hold null and can be used while binding
	and not required to use the pointer when it comes to smaller structs but using it here to get more used to with the pointers
	here,
	this is similar to  
	var formData *Details = &Details{}
	if you want to define non-go-idiomatic way
	*/
	fmt.Println("getting the form data..........")

	if err := c.BindJSON(&formData); err != nil {
		return nil, err
	}
	//-- adding && updating id as it is not coming from the form
	if formData.Type == "User" {
		formData.ID = len(userDetails) + 1
	} else {
		formData.ID = len(orgDetails) + 1
	}
	return formData, nil
}