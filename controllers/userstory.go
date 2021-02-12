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

// Operation about UserStory
type UserStoryController struct {
	beego.Controller
}

type AddUserStoryValidation struct {
	Title  string `valid:"Required; MaxSize(64)"`
	UserId int `valid:"Required"`
	StartDate string `valid:"Required"`
	EndDate string `valid:"Required"`
	IsPublish int `valid:"Required"`
}

// Custom Error Handler for the Validation of the StartDate and EndDate
func (u *AddUserStoryValidation) Valid(v *validation.Validation) {
	if len(u.StartDate) != 10 {
		v.SetError("StartDate", "Invalid start date format")
	}
	if len(u.EndDate) != 10 {
		v.SetError("EndDate", "Invalid end date format")
	}
}

// @Title create
// @Description add UserStory by the User
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserStoryController) Post() {
	var data AddUserStoryValidation
	// get the token from Header
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode:401, ErrorMessage:TokenError}
		this.ServeJSON()
	}
	// Unmarshal is Used to Convert into JSON
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	valid := validation.Validation{}
	story := AddUserStoryValidation{data.Title, getUser.UserId, data.StartDate,
		data.EndDate, data.IsPublish}
	b, _ := valid.Valid(&story)
	// validation does not pass .
	if !b {
		ErrMsg := make(map[string]string)
		for _, err := range valid.Errors {
			ErrMsg[err.Field] = err.Message
		}
		this.Data["json"] = ErrResponse{ErrCode: 400, ErrorMessage: ValidationError, Error:ErrMsg}
		this.ServeJSON()
	}
	StartDateObject := dateObject(data.StartDate)
	EndDateObject := dateObject(data.EndDate)

	// Raise Error When StartDate is invalid.
	if StartDateObject.Error != nil{
		dateError := StartDateObject.Error
		this.Data["json"] = ErrResponse{ErrCode: 400, ErrorMessage: ValidationError,
			Error:dateError.Error()}
		this.ServeJSON()
	}
	// Raise Error EndDate is invalid.
	if EndDateObject.Error != nil{
		dateError := EndDateObject.Error
		this.Data["json"] = ErrResponse{ErrCode: 400, ErrorMessage: ValidationError,
			Error:dateError.Error()}
		this.ServeJSON()
	}
	// Now Validate StartDate Should be less than EndDate
	if StartDateObject.Date.After(EndDateObject.Date){
		this.Data["json"] = ErrResponse{ErrCode: 400, ErrorMessage: ValidationError,
			Error:"StartDate should be less than EndDate."}
		this.ServeJSON()
	}
	userStory := models.UserStory{
		Title:data.Title,
		StartDate: StartDateObject.Date,
		EndDate: EndDateObject.Date,
		IsPublish: data.IsPublish,
	}
	response := models.AddUserStory(userStory, getUser.UserId)
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title get
// @Description get UserStory of the User
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [get]
func (this *UserStoryController) Get() {
	storyId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	// get the token from Header
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode:401, ErrorMessage:TokenError}
		this.ServeJSON()
	}
	response := models.GetUserStoryDetail(storyId, getUser.UserId)
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title delete
// @Description delete userStory by the User.
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [delete]
func (this *UserStoryController) Delete() {
	// Get the Id of the Preference from Input Param
	storyId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	// get the token from Header
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode: 401, ErrorMessage: TokenError}
		this.ServeJSON()
	}
	response := models.DeleteUserStory(storyId)
	this.Data["json"] = response
	this.ServeJSON()
}