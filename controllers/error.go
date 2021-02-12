
package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

//func (c *ErrorController) Error404() {
//	c.Data["json"] = Response{
//		code: 404,
//		message:  "Not Found",
//	}
//	c.ServeJSON()
//}
//func (c *ErrorController) Error401() {
//	c.Data["json"] = Response{
//		code: 401,
//		message:  "Permission denied",
//	}
//	c.ServeJSON()
//}
//func (c *ErrorController) Error403() {
//	c.Data["json"] = Response{
//		code: 403,
//		message:  "Forbidden",
//	}
//	c.ServeJSON()
//}