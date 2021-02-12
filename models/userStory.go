package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
	"time"
)

var (
	UserStorySuccess = "User story created successfully."
)

// UserStory model used to save the UserStory of the User.
type UserStory struct {
	Id int `orm:"pk;auto"`
	StartDate time.Time `orm:"auto_now_add;type(date)"`
	EndDate time.Time `orm:"type(date)"`
	Title string `orm:"size(64)"`
	IsPublish int `orm:"default(0)"`
	IsSaved int `orm:"default(0)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	User *User `orm:"rel(fk)"` // RelForeignKey relation
	UserStoryImage []*UserStoryImage `orm:"reverse(many)"` // reverse relationship of fk
}

// UserStoryImage used to store the Image of the UserStory.
type UserStoryImage struct {
	Id int `orm:"pk;auto"`
	StoryImage string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	CreatedBy int
	UpdatedBy int
	UserStory *UserStory `orm:"rel(fk)"` // RelForeignKey relation
}

func init(){
	orm.RegisterModel(new(UserStory))
	orm.RegisterModel(new(UserStoryImage))
}

// function used to add the UserStory
func AddUserStory(this UserStory, userId int)UserResponse{
	o := orm.NewOrm()
	userStory := new(UserStory)
	userStory.Title = this.Title
	userStory.StartDate = this.StartDate
	userStory.EndDate = this.EndDate
	userStory.IsPublish = this.IsPublish
	userStory.User = &User{Id: userId}
	_, err := o.Insert(userStory)
	if err == nil {
		return UserResponse{
			Code:    SuccessCode,
			Message: UserStorySuccess,
			Data:   userStory,
		}
	}
	return UserResponse{
		Code:    ErrorCode,
		Message: ErrorMessage,
		Data:   err,
	}
}

// function used to get the UserStory Details of the User
func GetUserStoryDetail(StoryId int, UserId int) UserResponse{
	var userStory[] *UserStory
	o := orm.NewOrm()
	// Get a QuerySetter object. preferences is table name
	qs := o.QueryTable("user_story")
	if StoryId == 0{
		qs.Filter("UserStory__id", UserId)
		qs.OrderBy("-UpdatedAt").Limit(10).All(&userStory, "Id", "Title",
			"StartDate", "EndDate", "IsPublish", "IsSaved", "CreatedAt", "UpdatedAt")
	}else{
		qs.Filter("Id", StoryId).All(&userStory, "Id", "Title", "StartDate", "EndDate",
			"IsPublish", "IsSaved", "CreatedAt", "UpdatedAt")
	}
	return UserResponse{
		Code:    SuccessCode,
		Message: GetMessage,
		Data:    userStory,
	}
}

// function used to delete the UserStory by the storyId.
func DeleteUserStory(StoryId int) UserResponse{
	o := orm.NewOrm()
	var userStory UserStory
	qs := o.QueryTable("user_story")
	err := qs.Filter("Id", StoryId).One(&userStory)
	if err == orm.ErrNoRows {
		// No result
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidId,
		}
	}
	// Now delete the user
	o.Delete(&userStory)
	return UserResponse{
		Code:    SuccessCode,
		Message: DeletedMessage,
	}
}