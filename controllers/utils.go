package controllers

import (
	"domio-api/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

type TokenResponse struct {
	Status  bool      `json:"status"`
	Admin bool         `json:"admin"`
	UserId  int      `json:"admin"`
}

func checkEmailExist(email string) bool{
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.User))
	// Now Filter With Email
	exist := qs.Filter("Email", email).Exist()
	return exist
}

func checkAuthToken(authToken string) bool{
	if authToken == ""{
		return true
	}else if len(authToken) <= 32 {
		return true
	}else {
		return false
	}
}

func verifyAuthToken(authToken string) TokenResponse{
	o := orm.NewOrm()
	var user models.User
	qs := o.QueryTable("user")
	err := qs.Filter("Token__Token", authToken).One(&user)
	fmt.Println("AUTH TOKEN QUERY", reflect.TypeOf(user.Status))
	if err == orm.ErrNoRows {
		// No result
		return TokenResponse{
			Status: false,
			Admin: false,
			UserId: 0,
		}
	}else if user.Status == 0 {
		return TokenResponse{
			Status: true,
			Admin: false,
			UserId: user.Id,
		}
	}else{
		return TokenResponse{
			Status: true,
			Admin: true,
			UserId: user.Id,
		}
	}
}

type DateResponse struct {
	Status  bool      `json:"status"`
	Date time.Time         `json:"date"`
	Error  error      `json:"error"`
}


// function used to convert string date to go lang date object
func dateObject(date string) DateResponse {
	dateObj, err := time.Parse(layoutISO, date)
	return DateResponse{
		Status: true,
		Date: dateObj,
		Error: err,
	}
}