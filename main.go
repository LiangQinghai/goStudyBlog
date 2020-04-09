package main

import (
	"encoding/gob"
	"github.com/LiangQinghai/goStudyBlog/models"
	_ "github.com/LiangQinghai/goStudyBlog/routers"
	_ "github.com/LiangQinghai/goStudyBlog/models"
	"github.com/astaxie/beego"
	"strings"
)

func main() {
	initSession()
	initTemplate()
	beego.Run()
}

func initSession(){
	gob.Register(&models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "HelloWorld"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func initTemplate(){
	beego.AddFuncMap("equrl", func(x,y string) bool{
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		return strings.Compare(s1, s2) == 0
	})
}