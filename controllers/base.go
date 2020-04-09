package controllers

import (
	"errors"
	"github.com/LiangQinghai/goStudyBlog/models"
	"github.com/astaxie/beego"
	"log"
)

const SESSION_USER_KEY  = "SESSION_USER_KEY"

type NestPrepare interface {
	NestsPrepare()
}

type BaseController struct {
	beego.Controller
	IsLogin bool
	User models.User
}


func (bc *BaseController) Prepare(){
	log.Println("BaseController")
	bc.IsLogin = false
	//Database.Using()
	us := bc.GetSession(SESSION_USER_KEY)
	if us != nil{
		if u,ok := us.(models.User); ok{
			bc.User = u
			bc.Data["User"] = u
			bc.IsLogin = true
		}
	}
	bc.Data["IsLogin"] = bc.IsLogin

	//如果controller实现了NestPrepare接口则执行
	bc.Data["Path"] = bc.Ctx.Request.RequestURI
	if app,ok := bc.AppController.(NestPrepare);ok{
		app.NestsPrepare()
	}
}

//校验数据
func (bc *BaseController) GetMustString(key, msg string) string {

	tmp := bc.GetString(key)

	if len(tmp) == 0 {
		bc.Abort500(errors.New("msg"))
	}

	return tmp
}

//抛出500错误
func (bc *BaseController) Abort500(err error){

	log.Print(err)
	bc.Data["error"] = err
	bc.Abort("500")

}