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
	if err = Config.DB.Save(b).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOneBook(b *Book, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Delete(b).Error; err != nil {
		return err
	}
	return nil
}
