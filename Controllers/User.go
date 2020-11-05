package Controllers

import (
	"fmt"
	"reflect"

	"github.com/Johanrim/example-web/Models"
	"github.com/Johanrim/example-web/Response"
	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUser(&user)
	if err != nil {
		Response.RespondJSON(c, 404, user)
	} else {
		Response.RespondJSON(c, 200, user)
	}
}

func CreateOneUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	//password

	err := Models.CreateOneUser(&user)
	if err != nil {
		Response.RespondJSON(c, 404, user)
	} else {
		Response.RespondJSON(c, 200, user)
	}
}

func GetOneUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetOneUser(&user, id)
	if err != nil {
		Response.RespondJSON(c, 404, user)
	} else {
		Response.RespondJSON(c, 200, user)
	}
}

func UpdateOneUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetOneUser(&user, id)
	if err != nil {
		Response.RespondJSON(c, 404, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateOneUser(&user, id)
	if err != nil {
		Response.RespondJSON(c, 404, user)
	} else {
		Response.RespondJSON(c, 200, user)
	}
}

func DeleteOneUser(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println(reflect.TypeOf(id))
	var user Models.User
	err := Models.DeleteOneUser(&user, id)
	if err != nil {
		Response.RespondJSON(c, 404, user)
	} else {
		Response.RespondJSON(c, 200, user)
	}
}
