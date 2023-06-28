package models

// https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:255;not null;" json:"password"`
	Phone     string `gorm:"size:255;" json:"phone"`
	Firstname string `gorm:"size:255;" json:"fristname"`
	Lastname  string `gorm:"size:255;" json:"lastname"`
	Wechat    string `gorm:"size:255;" json:"wechat"`
	Posts     []Post
}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
