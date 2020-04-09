package models

import "log"

//import "github.com/jinzhu/gorm"
//
//type User struct {
//	gorm.Model
//	Name string `gorm:"unique_index"`
//	Email string `gorm:"unique_index"`
//	Avatar string
//	Pwd string
//	Role int `gorm:"default:1"`
//}

//func (db *DB) QueryUserByEmailAndPassword(email, password string) (user User, err error){
//
//	return user,db.gdb.Model(&User{}).Where("email = ? and pwd = ?", email, password).Take(&user).Error
//
//}

type User struct {
	Id int `orm:"column(Id)"`
	UserName string `orm:"column(UserName)"`
	Password string `orm:"column(Password)"`
	Role int `orm:"column(Role)"`
}

func QueryUserByUserNameAndPassword(userName, password string) (user User, err error){

	user = User{UserName:userName, Password:password}

	err = Database.Read(&user)


	log.Print(user)

	return user, err

}

