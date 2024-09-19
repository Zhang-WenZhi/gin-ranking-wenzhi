package models

import (
	"time"
	"wenzhi.com/gin-ranking/dao"
)


type User struct {
	Id       int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	AddTime int64 `json:"addTime"`
	UpdateTime int64 `json:"updateTime"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func GetUserInfo(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func AddUser(username, password string) (int , error) {
	user := User{Password: password, Username: username, AddTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}



// import (
// 	"wenzhi.com/gin-ranking/dao"
// )
// type User struct {
// 	Id       int
// 	Username string
// }

// func (User) TableName() string {
// 	return "user"
// }

// func GetUserTest(id int) (User, error) {
// 	var user User
// 	err := dao.Db.Where("id =?", id).First(&user).Error
// 	return user, err
// }

// func GetUserListTest() ([]User, error) {
// 	var users []User
// 	err := dao.Db.Where("id < ?", 3).Find(&users).Error
// 	return users, err
// }

// func AddUser(username string) (int, error) {
// 	user := User{
// 		Username: username,
// 	}
// 	err := dao.Db.Create(&user).Error
// 	return user.Id, err
// }

// func UpdateUser(id int, username string) {
// 	dao.Db.Model(&User{}).Where("id =?", id).Update("username", username)
// }

// func DeleteUser(id int) error {
// 	err :=dao.Db.Delete(&User{}, "id =?", id).Error
// 	return err
// }