package Models

import (
	"fmt"

	"github.com/Johanrim/example-web/Config"
)

func GetAllUser(c *[]User) (err error) {
	if err = Config.DB.Find(c).Error; err != nil {
		return err
	}
	return nil
}

func CreateOneUser(c *User) (err error) {
	if err = Config.DB.Create(c).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUser(c *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(c).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOneUser(c *User, id string) (err error) {
	fmt.Println(c)
	if err = Config.DB.Save(c).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOneUser(c *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Delete(c).Error; err != nil {
		return err
	}
	return nil
}
