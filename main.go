package main

import (
	"fmt"

	"github.com/Johanrim/example-web/Config"
	"github.com/Johanrim/example-web/Models"
	"github.com/Johanrim/example-web/Routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("status: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})
	Config.DB.AutoMigrate(&Models.Book{}, &Models.Category{})
	Config.DB.AutoMigrate(&Models.Book{}).AddForeignKey("category_id", "category(id)", "CASCADE", "CASCADE")

	r := Routers.SetUpRouter()
	r.Run()
}
