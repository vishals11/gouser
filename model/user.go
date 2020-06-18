package model

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// User structure
type User struct {
	ID int `json:"id`

	Name     string `json:"name" validate:"required" gorm:"not null"`
	Email    string `json:"email" validate:"email,required" gorm:"unique;not null"`
	Password string `json:"password,omitempty" validate:"required" gorm:"not null"`
	PhoneNo  string `json:"phone_no"`

	Token string `json:"token" gorm:"not null"`
}

func CreateUser(user *User) error {
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPwd)

	_, err := GetUserFromEmail(user.Email)
	if err == nil {
		err = fmt.Errorf("User already exists")
		return err
	}

	err = GenerateToken(user)
	if err != nil {
		log.Println("Error Generating Token:", err)
		return err
	}

	err = db.Create(&user).Error
	if err != nil {
		log.Println("Error Creating user:", err)
		return err
	}

	// Dont return the hashed password to frontend
	user.Password = ""

	return nil
}

// Login structure
type Login struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

func LoginUser(login Login) (*User, error) {
	user, err := GetUserFromEmail(login.Email)
	if err != nil {
		log.Println("User not found")
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		err = fmt.Errorf("Password does not match")
		return nil, err
	}

	err = GenerateToken(user)
	if err != nil {
		log.Println("Error Generating Token:", err)
		return nil, err
	}

	err = db.Save(&user).Error
	if err != nil {
		log.Println("Error Updating user:", err)
		return nil, err
	}

	// Dont return the hashed password to frontend
	user.Password = ""

	return user, nil
}

func GetUserFromEmail(email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Println("Error searching user in database")
		return nil, err
	}

	return &user, nil
}
