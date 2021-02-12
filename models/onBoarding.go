package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
	"time"
)



// Create OnBoardingImage models
type OnBoardingImage struct {
	Id int `orm:"pk;auto"`
	ImageUrl string
	CreatedBy int
	UpdatedBy int
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

// create init function to register the models defined in this file.

func init() {
	orm.RegisterModel(new(OnBoardingImage))
}


// Function used to Create the OnBoarding Image by the SuperAdmin.
func CreateOnBoardingImage(ImageUrl string, UserId int) UserResponse {
	o := orm.NewOrm()
	var onBoard OnBoardingImage
	onBoard.ImageUrl = ImageUrl
	onBoard.CreatedBy = UserId
	onBoard.UpdatedBy = UserId
	o.Insert(&onBoard)
	return UserResponse{
		Code:    SuccessCode,
		Message: SaveMessage,
		Data:    onBoard,
	}
}

// Function used to get all onBoarding image.
func GetOnBoardingImage() UserResponse{
	var onBoard[] *OnBoardingImage
	o := orm.NewOrm()
	// Get a QuerySetter object. on_boarding_image is table name
	qs := o.QueryTable("on_boarding_image")
	qs.OrderBy("-UpdatedAt").Limit(10).All(&onBoard)
	// I have to discover how to use Offset in Limit
	return UserResponse{
		Code:    SuccessCode,
		Message: GetMessage,
		Data:    onBoard,
	}
}

// Function used to Update the OnBoarding Image by the Admin.
func UpdateOnBoardingImage(ImageUrl string, Id int, userId int) UserResponse {
	fmt.Println("ID is", Id)
	var onBoard OnBoardingImage
	o := orm.NewOrm()
	qs := o.QueryTable("on_boarding_image")
	err := qs.Filter("Id", Id).One(&onBoard)
	if err == orm.ErrNoRows {
		// No result
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidId,
		}
	}
	onBoard.ImageUrl = ImageUrl
	onBoard.UpdatedBy = userId
	_, err = o.Update(&onBoard)
	return UserResponse{
		Code:    SuccessCode,
		Message: UpdatedMessage,
		Data: onBoard,
	}
}

// Delete the OnBoardingImage by Id
func DeleteOnBoardingById(id int) UserResponse{
	o := orm.NewOrm()
	var onBoard OnBoardingImage
	qs := o.QueryTable("on_boarding_image")
	err := qs.Filter("Id", id).One(&onBoard)
	if err == orm.ErrNoRows {
		// No result
		return UserResponse{
			Code:    ErrorCode,
			Message: InvalidId,
		}
	}
	// Now delete the user
	o.Delete(&onBoard)
	return UserResponse{
		Code:    SuccessCode,
		Message: DeletedMessage,
	}
}