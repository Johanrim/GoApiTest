package Models

import (
	"fmt"

	"github.com/Johanrim/example-web/Config"
)

func GetAllCategory(c *[]Category) (err error) {
	if err = Config.DB.Find(c).Error; err != nil {
		return err
	}
	return nil
}

func CreateOneCategory(c *Category) (err error) {
	if err = Config.DB.Create(c).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCategory(c *Category, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(c).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOneCategory(c *Category, id string) (err error) {
	fmt.Println(c)
	if err = Config.DB.Save(c).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOneCategory(c *Category, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Delete(c).Error; err != nil {
		return err
	}
	return nil
}
