package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// Response Struct
type Response struct {
	SuccessCode int         `json:"code"`
	SuccessMessage  string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrResponse Struct
type ErrResponse struct {
	ErrCode int         `json:"code"`
	ErrorMessage  string      `json:"message"`
	Error  interface{} `json:"error"`
}