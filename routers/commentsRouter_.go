package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/LiangQinghai/goStudyBlog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/LiangQinghai/goStudyBlog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/LiangQinghai/goStudyBlog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/LiangQinghai/goStudyBlog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/LiangQinghai/goStudyBlog/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/LiangQinghai/goStudyBlog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetMessage",
            Router: `/message`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/LiangQinghai/goStudyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/LiangQinghai/goStudyBlog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/LiangQinghai/goStudyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/LiangQinghai/goStudyBlog/controllers:UserController"],
        beego.ControllerComments{
            Method: "LoginPage",
            Router: `/loginPage`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
