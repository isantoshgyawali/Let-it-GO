package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetFormData(c *gin.Context) (*Details, error) {
	fmt.Println()
	fmt.Println()
	//-- Parseform is used when traditional form data is sent over the http request
	//-- but as we are sending json using axios we will be using BindJSON
	//
	// if err := c.Request.ParseForm(); err != nil {
	// 	fmt.Println("Error Parsing the form data", err)
	// 	return nil, err
	// }
	// fmt.Println("Form Data: ", c.Request.PostForm)

	/**
	explicitly allocating so that it doesn't hold null and can be used while binding
	and here it is not required to use the pointer when it comes to smaller structs but using it here to get more used to with the pointers
	here,
	this is similar to this if you want to define non-go-idiomatic way
	var formData *Details = &Details{}
	*/
	formData:= &Details{}
	fmt.Println("getting the form data..........")

	if err := c.BindJSON(&formData); err != nil {
		return nil, err
	}
	userDetails = append(userDetails, formData)

	return formData, nil
}
