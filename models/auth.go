/*
 * @Author: tangqimin
 * @Date: 2021-11-12 11:17:15
 * @Description:
 * @LastEditTime: 2021-12-11 17:19:09
 * @LastEditors: tangqimin
 * @FilePath: \study-gin\models\auth.go
 */
package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	return auth.ID > 0
}
