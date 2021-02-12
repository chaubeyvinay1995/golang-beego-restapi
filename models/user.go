package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt" // for the encryption
	"time"
)

var (
	InvalidCredentials     = "Invalid Credentials"
	UserCreate = "User created successfully."
	LoggedInMessage    = "Logged in successfully."
	SaveMessage = "Record saved successfully."
	DeletedMessage = "Record deleted successfully."
	UpdatedMessage = "Record updated successFully."
	GetMessage = "Fetched data successfully."
	ErrorMessage = "Error in saving record."
	UniqueUserName = "Username is already acquired."
	PreferenceRemovedMessage = "Error in removing previous preferences."
	PreferenceRemovedSuccess = "Preference removed successfully."
	PreferenceAlreadyDeleted = "Preference is already deleted."
	PreferenceAddMessage = "Preferences added successfully."
	PreferenceErrorMessage = "Error while saving preferences."
	InvalidId = "Invalid id."
	InvalidPreferenceId = "Invalid preferences id."
	InvalidToken = "Invalid token."
	ErrorCode = 400
	SuccessCode = 200
)


// DeviceToken struct used to Store the DeviceToken.
type DeviceToken struct {
	Id int `orm:"pk;auto"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	DeviceToken string
	// Now define reverse relationship with the User Model With OneToOne field
	User *User `orm:"reverse(one)"`
}

// Token Struct used to store the token
type UserToken struct {
	Id int `orm:"pk;auto"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	Token string `orm:"unique"`
	// Now define reverse relationship with the User Model With OneToOne field
	User *User `orm:"reverse(one)"`
}

// User struct used to store the User info.
type User struct {
	Id       int `orm:"pk;auto"`
	Email string `orm:"unique"`
	Username string `orm:"null"`
	Password string
	UserImage string  `orm:"null"`
	Status int `orm:"default(0)"`
	SuperAdmin int `orm:"default(0)"`
	// One to One relationship with Token
	Token *UserToken `orm:"null;rel(one);on_delete(cascade)"`
	DeviceToken *DeviceToken `orm:"null; rel(one); on_delete(cascade)"`
	ForgotPasswordDetail *ForgotPassword `orm:"null;rel(one);on_delete(cascade)"`
	// Many to Many relationship with Preference
	Preferences []*Preference `orm:"rel(m2m)"`
	UserStory []*UserStory `orm:"reverse(many)"` // reverse relationship of fk
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

type Preference struct {
	Id int `orm:"pk;auto"`
	ImageUrl string
	Title string
	CreatedBy int
	UpdatedBy int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	//// Many to Many reverse relationship
	//Users []*User `orm:"reverse(many)"`
}

type ForgotPassword struct {
	Id int `orm:"pk;auto"`
	ForgotPasswordToken string `orm:"unique"`
	Status int `orm:"default(1)"`
	User *User `orm:"reverse(one)"`
	CreatedAt time.Time `orm:"auto_now_add:type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func init(){
	orm.RegisterModel(new(UserToken))
	orm.RegisterModel(new(DeviceToken))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Preference))
	orm.RegisterModel(new(ForgotPassword))
}

// Struct Used to Send the json response
type UserResponse struct {
	Code int         `json:"code"`
	Message  string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Function used to Store the Password in hash form
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

// Function Used to Check the Username exist or not.
func CheckUserNameExistOrNot(username string) bool {
	o := orm.NewOrm()
	var users User
	qs := o.QueryTable("user")
	err := qs.Filter("Username", username).One(&users)
	if err == orm.ErrNoRows {
		// No result
		return false
	}
	return true
}

// Func to Create the User
func CreateUser(user User) UserResponse {
	o := orm.NewOrm()
	users := new(User)
	users.Email = user.Email
	// Call hashPassword method to encrypt the password
	//users.Password = HashPassword(user.Password)
	users.Password = user.Password
	o.Insert(users)
	// Now create the token and associate with the User.
	token := new(UserToken)
	uuid := uuid.New()
	token.Token = uuid.String()
	tokenId,_  := o.Insert(token)
	o.QueryTable("user").Filter("Id", users.Id).Update(orm.Params{
		"Token":tokenId,
	})
	users.Token = token
	return UserResponse{
		Code:    SuccessCode,
		Message: UserCreate,
		Data: 			user,
	}

}

// Func to get the User Detail.
func GetUserDetails(email, password string) UserResponse{
	o := orm.NewOrm()
	var user User
	qs := o.QueryTable("user")
	err := qs.Filter("Email", email).Filter("Password", password).One(&user)
	if err == orm.ErrNoRows {
		// No result
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidCredentials,
		}
	}
	// To get the Token data
	o.LoadRelated(&user, "Token")
	// To get the User Preferences
	o.LoadRelated(&user, "Preferences")
	return UserResponse{
		Code:    SuccessCode,
		Message: LoggedInMessage,
		Data: 			user,
	}
}

// Get the User by Id
func GetUserById(authToken string) UserResponse{
	o := orm.NewOrm()
	var user User
	qs := o.QueryTable("user")
	err := qs.Filter("Token__Token", authToken).One(&user)
	if err == orm.ErrNoRows {
		// No result
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidId,
		}
	}
	// To get the Token data
	o.LoadRelated(&user, "Token")
	o.LoadRelated(&user, "Preferences")
	return UserResponse{
		Code:    SuccessCode,
		Message: GetMessage,
		Data: 			user,
	}
}

// Delete the User by Id
func DeleteUserById(authToken string) UserResponse{
	o := orm.NewOrm()
	var user User
	qs := o.QueryTable("user")
	err := qs.Filter("Token__Token", authToken).One(&user)
	if err == orm.ErrNoRows {
		// No result
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidId,
		}
	}
	// Now delete the user
	o.Delete(&user)
	return UserResponse{
		Code:    SuccessCode,
		Message: DeletedMessage,
	}
}

// Update the Use Detail by Id
func UpdateUserDetail(user User) UserResponse{
	o := orm.NewOrm()
	var users User
	qs := o.QueryTable("user")
	err := qs.Filter("Id", user.Id).One(&users)
	if err == orm.ErrNoRows {
		// No result
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidToken,
		}
	}
	if user.Password != "" {
		users.Password = HashPassword(user.Password)
	}
	if user.UserImage != ""{
		users.UserImage = user.UserImage
	}
	if user.Username != "" && users.Username != user.Username {
		usernameExist := CheckUserNameExistOrNot(user.Username)
		if usernameExist{
			return UserResponse{
				Code:    ErrorCode,
				Message: UniqueUserName,
				Data:           nil,
			}
		}
		users.Username = user.Username
	}
	
	if user.Status == 0 || user.Status == 1{
		users.Status = user.Status
	}
	_, err = o.Update(&users)
	if err == nil{
		// Now Load the Related token
		o.LoadRelated(&users, "Token")
		return UserResponse{
			Code:    SuccessCode,
			Message: UpdatedMessage,
			Data:           users,
		}
	}
	return UserResponse{
		Code:    ErrorCode,
		Message: ErrorMessage,
		Data:           nil,
	}
}

// function used to create the preference
func CreatePreference(this Preference, UserId int) UserResponse{
	o := orm.NewOrm()
	preference := new(Preference)
	preference.Title = this.Title
	preference.ImageUrl = this.ImageUrl
	preference.CreatedBy = UserId
	preference.UpdatedBy = UserId
	_, err := o.Insert(preference)
	if err == nil {
		return UserResponse{
			Code:    SuccessCode,
			Message: SaveMessage,
			Data:   preference,
		}
	}
	return UserResponse{
		Code:    ErrorCode,
		Message: ErrorMessage,
		Data:   err,
	}
}

// function used to get the preferences
func GetPreference(Id int) UserResponse{
	var preferences[] *Preference
	o := orm.NewOrm()
	// Get a QuerySetter object. preferences is table name
	qs := o.QueryTable("preference")
	if Id == 0{
		qs.OrderBy("-UpdatedAt").Limit(10).All(&preferences,
			"Title", "ImageUrl", "Id", "CreatedAt", "UpdatedAt")
	}else{
		qs.Filter("Id", Id).All(&preferences)
	}
	return UserResponse{
		Code:    SuccessCode,
		Message: GetMessage,
		Data:    preferences,
	}
}

// function used to delete the Preference
func DeletePreference(id int)UserResponse{
	o := orm.NewOrm()
	var preference Preference
	qs := o.QueryTable("preference")
	err := qs.Filter("Id", id).One(&preference)
	if err == orm.ErrNoRows {
		// No result
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidId,
		}
	}
	// Now delete the user
	o.Delete(&preference)
	return UserResponse{
		Code:    SuccessCode,
		Message: DeletedMessage,
	}
}

// function used to Update the Preference
func UpdatePreference(this Preference, UserId int) UserResponse {
	o := orm.NewOrm()
	var preference Preference
	qs := o.QueryTable("preference")
	err := qs.Filter("Id", this.Id).One(&preference)
	if err == orm.ErrNoRows {
		// No result
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidId,
		}
	}
	preference.ImageUrl = this.ImageUrl
	preference.Title = this.Title
	preference.UpdatedBy = UserId
	_, err = o.Update(&preference)
	if err == nil{
		// Now Load the Related token
		return UserResponse{
			Code:    SuccessCode,
			Message: UpdatedMessage,
			Data:           preference,
		}
	}
	return UserResponse{
		Code:    ErrorCode,
		Message: ErrorMessage,
		Data:           nil,
	}

}

// function used to add the preferences of the user
func AddUserPreference(preferences [3]int, size int,  UserId int) UserResponse{
	var preference[] Preference
	o := orm.NewOrm()
	qs := o.QueryTable("preference")
	count, err := qs.Filter("Id__in", preferences).All(&preference)
	if err == orm.ErrNoRows {
		// No result
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidPreferenceId,
		}
	}else if count < 3{
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidPreferenceId,
		}
	}
	// Now associate the preferences with UserId
	user := User{Id: UserId}
	m2m := o.QueryM2M(&user, "Preferences")
	// Clear the Previous Associated Preferences.
	_, removeErr := m2m.Clear()
	if removeErr != nil {
		return UserResponse{
			Code:    ErrorCode,
			Message: PreferenceRemovedMessage,
			Data:  nil,
		}
	}
	//Now add the new Preferences
	_, err = m2m.Add(preferences)
	if err == nil {
		return UserResponse{
			Code:    SuccessCode,
			Message: PreferenceAddMessage,
			Data:   preferences,
		}
	}
	return UserResponse{
		Code:    ErrorCode,
		Message: PreferenceErrorMessage,
		Data:   nil,
	}
}

// function used to get the user preferences
func GetUserPreferences(authToken string) UserResponse{
	o := orm.NewOrm()
	var user User
	qs := o.QueryTable("user")
	err := qs.Filter("Token__Token", authToken).One(&user, "Id", "Email",
		"Status", "CreatedAt", "UpdatedAt")
	if err == orm.ErrNoRows {
		// No result
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidId,
		}
	}
	// To get the Token data
	o.LoadRelated(&user, "Token")
	o.LoadRelated(&user, "Preferences")
	return UserResponse{
		Code:    SuccessCode,
		Message: GetMessage,
		Data: 			user,
	}
}

// function used to delete the user preferences
func DeleteUserPreferences(userId int) UserResponse{
	o := orm.NewOrm()
	user := User{Id: userId}
	m2m := o.QueryM2M(&user, "Preferences")
	// Clear the Previous Associated Preferences.
	count, removeErr := m2m.Clear()
	if removeErr != nil {
		return UserResponse{
			Code:    ErrorCode,
			Message: PreferenceRemovedMessage,
			Data:  nil,
		}
	}else if count == 0{
		return UserResponse{
			Code:    ErrorCode,
			Message: PreferenceAlreadyDeleted,
			Data:  nil,
		}
	}
	return UserResponse{
		Code:    SuccessCode,
		Message: PreferenceRemovedSuccess,
		Data:  nil,
	}
}