package Controllers

import (
	"fmt"

	"github.com/Johanrim/example-web/Models"
	"github.com/Johanrim/example-web/Response"
	"github.com/gin-gonic/gin"
)

func ListBook(c *gin.Context) {
	var book []Models.Book
	err := Models.GetAllBook(&book)
	if err != nil {
		Response.ErrorRespondJSON(c, 404, err)
	} else {
		Response.RespondJSON(c, 200, book)
	}
}

func CreateOneBook(c *gin.Context) {
	var book Models.Book
	c.BindJSON(&book)
	err := Models.CreateOneBook(&book)
	if err != nil {
		Response.ErrorRespondJSON(c, 404, err)
	} else {
		Response.RespondJSON(c, 200, book)
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetOneBook(&book, id)
	if err != nil {
		Response.ErrorRespondJSON(c, 404, err)
	} else {
		fmt.Println(book)
		Response.RespondJSON(c, 200, book)
	}
}

func UpdateOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetOneBook(&book, id)
	if err != nil {
		Response.ErrorRespondJSON(c, 404, err)
	}
	c.BindJSON(&book)
	err = Models.UpdateOneBook(&book, id)
	if err != nil {
		Response.ErrorRespondJSON(c, 404, err)
	} else {
		Response.RespondJSON(c, 200, book)
	}
}

func DeleteOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.DeleteOneBook(&book, id)
	if err != nil {
		Response.ErrorRespondJSON(c, 404, err)
	} else {
		Response.RespondJSON(c, 200, book)
	}
}
