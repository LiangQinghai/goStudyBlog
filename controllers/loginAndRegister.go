package controllers

import (
	"github.com/LiangQinghai/goStudyBlog/models"
	"github.com/astaxie/beego/orm"
)

type LoginAndRegisterController struct {
	BaseController
}

//登录页
// @router /loginPage [get]
func (lac * LoginAndRegisterController) LoginPage(){
	lac.TplName = "login.html"
}
//注册页
//@router /regPage [get]
func (lac *LoginAndRegisterController) RegPage(){
	lac.TplName = "register.html"
}


//登录
//@router /login [post]
func (lac *LoginAndRegisterController) Login(){

	// handle something and then redirect to index Page
	userName := lac.GetString("userName")
	password := lac.GetString("password")
	var user models.User
	err := orm.NewOrm().QueryTable("User").Filter("user_name", userName).One(&user)
	if err != nil{
		lac.Data["error"] = err
		lac.Abort("500")
	}else if user.Password != password {
		lac.Data["error"] = "用户名或密码错误。。。"
		lac.Abort("500")
	}else {

		lac.SetSession("user", &user)
		//登录成功
		lac.Redirect("/", 302)
	}

}

//注册
//@router /register [post]
func (lac *LoginAndRegisterController) Register(){

	userName := lac.GetString("userName")
	password := lac.GetString("password")

	//validate
	if userName == "" || len(userName) == 0 || len(userName) < 3 || len(userName) > 20 || password == "" || len(password) == 0 || len(password) < 6 || len(password) > 12 {
		lac.Data["error"] = "用户名或密码格式不正确"
		lac.Abort("500")
		return
	}

	ormer := orm.NewOrm()
	var user models.User
	if err := ormer.QueryTable("user").Filter("user_name", userName).One(&user);err == nil{
		lac.Data["error"] = "用户名已经存在"
		lac.Abort("500")
		return
	}

	user = models.User{UserName:userName,Password:password,Role:1}
	ormer.Insert(&user)
	//注册成功
	lac.Redirect("/login", 302)
}