package controllers

import "github.com/LiangQinghai/goStudyBlog/models"

type UserController struct {
	BaseController
}

// @router /login [post]
func (uc *UserController) Login()  {

	userName := uc.GetMustString("userName", "Email Must Not Be Empty")
	pwd := uc.GetMustString("password", "Password Must Not Be Empty")

	var (
		user models.User
		err error
	)

	if user,err = models.QueryUserByUserNameAndPassword(userName, pwd); err != nil{
		uc.Abort500(err)
	}

	uc.SetSession(SESSION_USER_KEY, user)

	uc.Redirect("/", 200)

}

// @router /loginPage [get]
func (uc *UserController) LoginPage() {
	uc.TplName = "user.html"
}