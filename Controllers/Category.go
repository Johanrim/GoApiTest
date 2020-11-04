package Controllers

import (
	"fmt"
	"reflect"

	"github.com/Johanrim/example-web/Models"
	"github.com/Johanrim/example-web/Response"
	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	var category []Models.Category
	err := Models.GetAllCategory(&category)
	if err != nil {
		Response.RespondJSON(c, 404, category)
	} else {
		Response.RespondJSON(c, 200, category)
	}
}

func CreateOneCategory(c *gin.Context) {
	var category Models.Category
	c.BindJSON(&category)
	err := Models.CreateOneCategory(&category)
	if err != nil {
		Response.RespondJSON(c, 404, category)
	} else {
		Response.RespondJSON(c, 200, category)
	}
}

func GetOneCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	var category Models.Category
	err := Models.GetOneCategory(&category, id)
	if err != nil {
		Response.RespondJSON(c, 404, category)
	} else {
		Response.RespondJSON(c, 200, category)
	}
}

func UpdateOneCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	var category Models.Category
	err := Models.GetOneCategory(&category, id)
	if err != nil {
		Response.RespondJSON(c, 404, category)
	}
	c.BindJSON(&category)
	err = Models.UpdateOneCategory(&category, id)
	if err != nil {
		Response.RespondJSON(c, 404, category)
	} else {
		Response.RespondJSON(c, 200, category)
	}
}

func DeleteOneCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println(reflect.TypeOf(id))
	var category Models.Category
	err := Models.DeleteOneCategory(&category, id)
	if err != nil {
		Response.RespondJSON(c, 404, category)
	} else {
		Response.RespondJSON(c, 200, category)
	}
}
