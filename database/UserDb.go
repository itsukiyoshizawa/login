package db

import (
	"main/data"
)

//ユーザIdとパスワードをもとにユーザの情報を取得する
func UserSelect(userId string, password string) data.User {
	d := GormConnect()
	selData := data.User{}
	d.First(&selData, "user_id=? and password = ?", userId, password)
	defer d.Close()
	return selData
}
