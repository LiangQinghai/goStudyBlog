package controllers

type IndexController struct {
	BaseController
}

// @router / [get]
func (ic *IndexController) Get()  {
	ic.TplName = "index.html"
}

// @router /message [get]
func (ic *IndexController) GetMessage(){
	ic.TplName = "message.html"
}


// @router /about [get]
func (ic *IndexController) GetAbout(){
	ic.TplName = "about.html"
}