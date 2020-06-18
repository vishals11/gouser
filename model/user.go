package model

import "golang.org/x/crypto/bcrypt"

// User structure
type User struct {
	ID int `json:"id`

	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password,omitempty" validate:"required"`
}

func CreateUser(user *User) (*User, error) {
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPwd)

	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// Dont return the hashed password to frontend
	user.Password = ""

	return user, nil
}
