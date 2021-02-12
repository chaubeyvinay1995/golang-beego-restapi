// Controller file for the User Preferences.
package controllers

import (
	"domio-api/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strconv"
)

// Preference Struct to handle the Preferences.
type PreferenceController struct {
	beego.Controller
}

// UserPreferenceController Struct to handle the Preferences.
type UserPreferenceController struct {
	beego.Controller
}

// Define Struct For the Preference Validation while creating preferences
type AddPreferenceValidation struct {
	ImageUrl  string `valid:"Required"`
	Title string `valid:"MinSize(4); MaxSize(16)"`
}

// Struct Used for the Validation While Updating the Preference
type UpdatePreferenceValidation struct {
	ImageUrl  string `valid:"Required"`
	Title string `valid:"MinSize(4); MaxSize(16)"`
	Id int `valid:"Required"`
}

// Struct used for the Validation of UserPreference
type AddUserPreferenceValidation struct {
	Preferences [3]int `valid:"Required"`
}


// @Title create
// @Description Create Preference by the Admin.
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *PreferenceController) Post(){
	var data AddPreferenceValidation
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
		this.Data["json"] = ErrResponse{ErrCode:407, ErrorMessage:SuperAdmin}
		this.ServeJSON()
	}
	// Unmarshal is Used to Convert into JSON
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	valid := validation.Validation{}
	userData := AddPreferenceValidation{data.ImageUrl, data.Title}
	b, _ := valid.Valid(&userData)
	// validation does not pass .
	if !b {
		ErrMsg := make(map[string]string)
		for _, err := range valid.Errors {
			ErrMsg[err.Field] = err.Message
		}
		this.Data["json"] = ErrResponse{ErrCode: 400, ErrorMessage: ValidationError, Error:ErrMsg}
		this.ServeJSON()
	}
	//Now create the Preferences Struct
	preference := models.Preference{
		Title:             data.Title,
		ImageUrl:               data.ImageUrl,
	}
	response := models.CreatePreference(preference, getUser.UserId)
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title get
// @Description get Preference by the Admin/User.
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [get]
func (this *PreferenceController) Get(){
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	// get the token from Header
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode:401, ErrorMessage:TokenError}
		this.ServeJSON()
	}
	response := models.GetPreference(id)
	this.Data["json"] = response
	this.ServeJSON()
}


// @Title delete
// @Description delete Preference by the Admin.
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [delete]
func (this *PreferenceController) Delete(){
	// Get the Id of the Preference from Input Param
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
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
		this.Data["json"] = ErrResponse{ErrCode:407, ErrorMessage:SuperAdmin}
		this.ServeJSON()
	}
	response := models.DeletePreference(id)
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title update
// @Description update user preference
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [put]
func (this *PreferenceController) Put(){
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
		this.Data["json"] = ErrResponse{ErrCode:407, ErrorMessage:SuperAdmin}
		this.ServeJSON()
	}
	var data UpdatePreferenceValidation
	// Unmarshal is Used to Convert into JSON
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	valid := validation.Validation{}
	userData := UpdatePreferenceValidation{data.ImageUrl, data.Title, data.Id}
	b, _ := valid.Valid(&userData)
	// validation does not pass .
	if !b {
		ErrMsg := make(map[string]string)
		for _, err := range valid.Errors {
			ErrMsg[err.Field] = err.Message
		}
		this.Data["json"] = ErrResponse{ErrCode: 400, ErrorMessage: ValidationError, Error:ErrMsg}
		this.ServeJSON()
	}
	preference := models.Preference{
		Id:                   data.Id,
		Title:             data.Title,
		ImageUrl:               data.ImageUrl,
	}
	response := models.UpdatePreference(preference, getUser.UserId)
	this.Data["json"] = response
	this.ServeJSON()
}

// Custom Validation for the UserPreferences
func (u *AddUserPreferenceValidation) Valid(v *validation.Validation) {
	for _, preference := range u.Preferences{
		if preference == 0{
			v.SetError("Preferences", "Minimum 3 preferences is required.")
		}
	}
}

// @Title create
// @Description Create UserPreference by the User.
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserPreferenceController) Post() {
	var data AddUserPreferenceValidation
	// get the token from Header
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode: 401, ErrorMessage: TokenError}
		this.ServeJSON()
	}
	// Unmarshal is Used to Convert into JSON
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	valid := validation.Validation{}
	userData := AddUserPreferenceValidation{data.Preferences}
	b, _ := valid.Valid(&userData)
	// validation does not pass .
	if !b {
		ErrMsg := make(map[string]string)
		for _, err := range valid.Errors {
			ErrMsg[err.Field] = err.Message
		}
		this.Data["json"] = ErrResponse{ErrCode: 400, ErrorMessage: ValidationError, Error:ErrMsg}
		this.ServeJSON()
	}
	response := models.AddUserPreference(data.Preferences, 3, getUser.UserId)
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title get
// @Description get UserPreference by the User.
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [get]
func (this *UserPreferenceController) Get() {
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode: 401, ErrorMessage: TokenError}
		this.ServeJSON()
	}
	// Now call the Function to get the UserPreference Details
	response := models.GetUserPreferences(authToken)
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title delete
// @Description Delete UserPreference by the User.
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [delete]
func (this *UserPreferenceController) Delete(){
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode: 401, ErrorMessage: TokenError}
		this.ServeJSON()
	}
	// Now call the Function to Delete the UserPreference Details
	response := models.DeleteUserPreferences(getUser.UserId)
	this.Data["json"] = response
	this.ServeJSON()
}