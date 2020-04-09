package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//type DB struct {
//	gdb *gorm.DB
//}
//
//func (db *DB) Begin(){
//	db.gdb = db.gdb.Begin()
//}
//
//func (db *DB) RollBack(){
//	db.gdb = db.gdb.Rollback()
//}
//
//func (db *DB) Commint()  {
//	db.gdb = db.gdb.Commit()
//}
//
//var (
//	db *gorm.DB
//)

//func NewDB() *DB{
//	return &DB{gdb:db}
//}
//
//func init(){
//	if err := os.Mkdir("data",0777);err != nil{
//		panic("Failed to connect to database..." + err.Error())
//	}
//
//	db, err := gorm.Open("sqlite3", "data/data.db")
//
//	if err != nil {
//		panic("Failed to connect..." + err.Error())
//	}
//
//	//同步表结构
//	db.AutoMigrate(&User{})
//
//	var count int
//	if err := db.Model(&User{}).Count(&count).Error; err != nil && count == 0{
//		db.Create(&User{
//			Name: "admin",
//			Email: "liangqinghai@live.com",
//			Pwd: "123456",
//			Avatar: "/static/images/info-img.png",
//			Role: 0,
//		})
//	}
//
//}


var (
	Database orm.Ormer
)

func init() {
	orm.RegisterModel(new(User))
	er1 := orm.RegisterDriver("mysql", orm.DRMySQL)
	if er1 != nil {
		panic(er1)
	}
	er2 := orm.RegisterDataBase("default", "mysql", "root:liang@tcp(localhost:3306)/goblog?charset=utf8")
	if er2 != nil {
		panic(er2)
	}

	Database = orm.NewOrm()
}