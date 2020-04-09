package routers

import (
	"github.com/LiangQinghai/goStudyBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.IndexController{})

    //错误控制器
    beego.ErrorController(&controllers.ErrorController{})
    //配置注解路由
    beego.Include(&controllers.IndexController{}, &controllers.UserController{})
}
