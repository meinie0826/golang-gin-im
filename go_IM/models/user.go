package models

import (
	"fmt"
	"golangim/utils"

	"gorm.io/gorm"
)

var (
	ID int
)

type User struct {
	Id         int
	Name       string
	Password   string
	Repassword string
}

func (table *User) TableName() string {
	return "user"
}

func GetUserList() []*User {
	data := make([]*User, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func FindUserByName(name string) User {
	user := User{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}
func DeleteUser(user User) *gorm.DB {
	return utils.DB.Delete(&user)
}
func UpdateUser(user User) *gorm.DB {
	//db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	return utils.DB.Model(&User{}).Where("name=?", user.Name).Updates(User{Name: user.Name, Password: user.Password, Repassword: user.Repassword})
}
