package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.TplName = "errors/404.html"
}

func (c *ErrorController) Error500() {
	c.TplName = "errors/500.html"
}

func (c *ErrorController) ErrorDb() {
	c.TplName = "errors/dberror.html"
}