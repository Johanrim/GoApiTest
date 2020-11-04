package Routers

import (
	"github.com/Johanrim/example-web/Controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.GET("book", Controllers.ListBook)
		v1.POST("book", Controllers.CreateOneBook)
		v1.GET("book/:id", Controllers.GetOneBook)
		v1.PUT("book/:id", Controllers.UpdateOneBook)
		v1.DELETE("book/:id", Controllers.DeleteOneBook)
	}
	v2 := r.Group("/")
	{
		v2.GET("category", Controllers.ListCategory)
		v2.POST("category", Controllers.CreateOneCategory)
		v2.GET("category/:id", Controllers.GetOneCategory)
		v2.PUT("category/:id", Controllers.UpdateOneCategory)
		v2.DELETE("category/:id", Controllers.DeleteOneCategory)
	}

	// v2 := r.Group("/author/")
	// {
	// 	v2.GET("", Controller.)
	// }
	return r
}
