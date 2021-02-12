package controllers

import (
	"domio-api/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strconv"
)

func init() {
	validation.SetDefaultMessage(map[string]string{
		"Required":     "This field is required.",
	})
}

// Operation about OnBoarding
type OnBoardingController struct {
	beego.Controller
}

type AdminOnBoardingController struct {
	beego.Controller
}

type OnBoardingImageValidation struct {
	ImageUrl  string `valid:"Required"`
	Id int `valid:"Required"`
}

type UpdateOnBoardingImageValidation struct {
	ImageUrl  string `valid:"Required"`
	Id int `valid:"Required"`
	UserId int `valid:"Required"`
}

// @Title create
// @Description add onBoarding Image by Admin
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *OnBoardingController) Post() {
	var data OnBoardingImageValidation
	// get the token from Header
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode:401, ErrorMessage:TokenError}
		this.ServeJSON()
	}
	// Check user is Admin or Not
	if !getUser.Admin {
		this.Data["json"] = ErrResponse{ErrCode:407, ErrorMessage:"User is not admin."}
		this.ServeJSON()
	}
	// Unmarshal is Used to Convert into JSON
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	valid := validation.Validation{}
	onBoard := OnBoardingImageValidation{data.ImageUrl, getUser.UserId}
	b, _ := valid.Valid(&onBoard)
	// validation does not pass .
	if !b {
		ErrMsg := make(map[string]string)
		for _, err := range valid.Errors {
			ErrMsg[err.Field] = err.Message
		}
		this.Data["json"] = ErrResponse{ErrCode: 400, ErrorMessage: ValidationError, Error:ErrMsg}
		this.ServeJSON()
	}
	response := models.CreateOnBoardingImage(data.ImageUrl,  getUser.UserId)
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title get
// @Description get the OnBoarding Image
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [get]
func (this *OnBoardingController) Get() {
	response := models.GetOnBoardingImage()
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title update
// @Description update the onBoarding image by Admin
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [patch]
func (this *AdminOnBoardingController) Update() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	var data UpdateOnBoardingImageValidation
	// Unmarshal is Used to Convert into JSON
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode:401, ErrorMessage:TokenError}
		this.ServeJSON()
	}
	// Check user is Admin or Not
	if !getUser.Admin {
		this.Data["json"] = ErrResponse{ErrCode:400, ErrorMessage:SuperAdmin}
		this.ServeJSON()
	}

	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	valid := validation.Validation{}
	onBoard := UpdateOnBoardingImageValidation{data.ImageUrl, id, getUser.UserId}
	b, _ := valid.Valid(&onBoard)
	// validation does not pass .
	if !b {
		ErrMsg := make(map[string]string)
		for _, err := range valid.Errors {
			ErrMsg[err.Field] = err.Message
		}
		this.Data["json"] = ErrResponse{ErrCode: 400, ErrorMessage: ValidationError, Error: ErrMsg}
		this.ServeJSON()
	}
	response := models.UpdateOnBoardingImage(data.ImageUrl, id, getUser.UserId)
	this.Data["json"] = response
	this.ServeJSON()
}


//@Title delete
//@Description delete the onBoarding image by Admin
//@Param   body        body    models.Object   true        "The object content"
//@Success 200 {string} models.Object.Id
//@Failure 403 body is empty
//@router / [delete]
func (this *AdminOnBoardingController) Delete(){
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode:401, ErrorMessage:TokenError}
		this.ServeJSON()
	}
	//Check user is Admin or Not
	if !getUser.Admin {
		this.Data["json"] = ErrResponse{ErrCode:400, ErrorMessage:SuperAdmin}
		this.ServeJSON()
	}
	response := models.DeleteOnBoardingById(id)
	this.Data["json"] = response
	this.ServeJSON()
}