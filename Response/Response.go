package Response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	Meta   interface{}
	Data   interface{}
}

type ErrorRespondData struct {
	Status  int
	Message string
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	fmt.Println("status ", status)
	var res ResponseData
	res.Status = status
	res.Data = payload
	w.JSON(200, res)
}

func ErrorRespondJSON(w *gin.Context, status int, err error) {
	fmt.Println("status: ", status)
	var res ErrorRespondData
	res.Status = status
	res.Message = err.Error()
	w.JSON(200, res)
}
