
package controllers

import (
	"domio-api/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strings"
)

func init(){
	validation.SetDefaultMessage(map[string]string{
		"Required": "This field is required.",
		"Email": "Must be in the correct email format.",
		"Min": "The minimum allowed value is% d",
		"Max": "The maximum allowed value is% d",
		"Range": "Must be in the range of% d to% d",
	})
}

var (
	EmailIsRegistered     = "Email is already Registered."
	ValidationError = "Validation Error."
	TokenError = "Invalid Token."
	SuperAdmin = "User is not admin."

)

// Initiate the Controller
type RegisterController struct {
	beego.Controller
}

type UserValidation struct {
	Email  string `valid:"Email; MaxSize(32)"` // Need to be a valid Email address and no more than 100 characters.
	Password    string    `valid:"MinSize(8); MaxSize(16)"` // 8 <= Password <= 16, only valid in this range
}

// Custom Error Handling
func (u *UserValidation) Valid(v *validation.Validation) {
	if strings.HasPrefix(u.Password, "admin") {
		// Set error messages of Name by SetError and HasErrors will return true
		v.SetError("Password", "Can't use admin in Password")
	}
}
// @Title create
// @Description register user object
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *RegisterController) Post(){
	var data UserValidation
	// Unmarshal is Used to Convert into JSON
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	valid := validation.Validation{}
	register := UserValidation{data.Email, data.Password}
	b, _ := valid.Valid(&register)
	// validation does not pass .
	if !b {
		ErrMsg := make(map[string]string)
		for _, err := range valid.Errors {
			ErrMsg[err.Field] = err.Message
		}
		this.Data["json"] = ErrResponse{ErrCode:400, ErrorMessage:ValidationError, Error:ErrMsg}
		this.ServeJSON()
	}
	// Now check that Given email exist in System or not.
	emailExist := checkEmailExist(data.Email)
	if emailExist{
		this.Data["json"] = ErrResponse{ErrCode:400, ErrorMessage:EmailIsRegistered}
		this.ServeJSON()
	}else {
		user := models.User{
			Email: data.Email,
			Password: data.Password,
		}
		response := models.CreateUser(user)
		this.Data["json"] = response
		this.ServeJSON()
	}
}
// Initiate the Controller
type LoginController struct {
	beego.Controller
}

// @Title create
// @Description Login the user
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *LoginController) Post(){
	var data UserValidation
	// Unmarshal is Used to Convert into JSON
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	valid := validation.Validation{}
	userData := UserValidation{data.Email, data.Password}
	b, _ := valid.Valid(&userData)
	// validation does not pass .
	if !b {
		ErrMsg := make(map[string]string)
		for _, err := range valid.Errors {
			ErrMsg[err.Field] = err.Message
		}
		this.Data["json"] = ErrResponse{ErrCode:400, ErrorMessage:ValidationError, Error:ErrMsg}
		this.ServeJSON()
	}
	// Now get the user
	response := models.GetUserDetails(userData.Email, userData.Password)
	this.Data["json"] = response
	this.ServeJSON()
}


// Initiate the Controller
type UserDetailController struct {
	beego.Controller
}
// @Title create
// @Description Get user details
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [get]
func (this *UserDetailController) Get(){
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode:401, ErrorMessage:TokenError}
		this.ServeJSON()
	}
	response := models.GetUserById(authToken)
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title delete
// @Description delete the user
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [delete]
func (this *UserDetailController) Delete(){
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode:401, ErrorMessage:TokenError}
		this.ServeJSON()
	}
	response := models.DeleteUserById(authToken)
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title update
// @Description update user detail
// @Param   body        body    models.Object   true        "The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [patch]
func (this *UserDetailController) Update(){
	authToken := this.Ctx.Request.Header.Get("Authorization")
	// Call Function to verify token
	getUser := verifyAuthToken(authToken)
	// Check the token is valid or not
	if getUser.Status == false {
		this.Data["json"] = ErrResponse{ErrCode:401, ErrorMessage:TokenError}
		this.ServeJSON()
	}
	var data models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	user := models.User{
		Id:                   getUser.UserId,
		Username:             data.Username,
		Status:               data.Status,
		UserImage: data.UserImage,
	}
	response := models.UpdateUserDetail(user)
	this.Data["json"] = response
	this.ServeJSON()
}