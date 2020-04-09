package controllers

type ErrorController struct {
	BaseController
}

func (ec *ErrorController) Error404()  {
	ec.TplName = "404.html"
}
