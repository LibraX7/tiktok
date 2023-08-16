package dao

import (
	"fmt"
	"go_gin/global"
	"go_gin/models"
)

var user models.User

func UserLogin(username string, password string) (*models.User, bool) {
	//调用DB
	rows := global.DB.Where("name = ? AND password = ?", username, password).Find(&user)
	fmt.Println(&user)
	//查询失败
	if rows.RowsAffected < 1 {
		return &user, false
	}
	return &user, true
}
func UserCreate(user *models.User) (*models.User, bool) {
	//调用DB
	rows := global.DB.Create(user)
	fmt.Println(user)
	//查询失败
	if rows.RowsAffected < 1 {
		return user, false
	}
	return user, true
}
