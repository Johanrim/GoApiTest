package Models

import (
	"fmt"

	"github.com/Johanrim/example-web/Config"
)

func GetAllBook(b *[]Book) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func CreateOneBook(b *Book) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	fmt.Println(b)
	return nil
}

func GetOneBook(b *Book, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOneBook(b *Book, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteOneBook(b *Book, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
