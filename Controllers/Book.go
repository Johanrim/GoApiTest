package Controllers

import (
	"fmt"
	"reflect"

	"github.com/Johanrim/example-web/ApiHelpers"
	"github.com/Johanrim/example-web/Models"
	"github.com/gin-gonic/gin"
)

func ListBook(c *gin.Context) {
	var book []Models.Book
	err := Models.GetAllBook(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func CreateOneBook(c *gin.Context) {
	var book Models.Book
	c.BindJSON(&book)
	err := Models.CreateOneBook(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func UpdateOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	}
	c.BindJSON(&book)
	err = Models.UpdateOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func DeleteOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println(reflect.TypeOf(id))
	var book Models.Book
	err := Models.DeleteOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}
