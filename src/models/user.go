package models

import (
	"golang.org/x/crypto/bcrypt"
)

//User struct to create user table yn DB with automigration
type User struct {
	Model
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"-"` //not returning in DB user data query
	IsAmbassador bool   `json:"-"`
}

//PasswordEncryption receiver encrypts and set the user pass  field using bcrypt library
func (user *User) SetAndEncryptPassword(pass string) {
	//number of layer for encryption algo
	cost := 8
	//GeneratesFormPassword only accepts a slice of bytes []byte
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), cost)
	//set in Password field the encrypted password as string
	user.Password = string(bytes)
}

//comparePassword receiver checks user hashed pass in DB wit the user pass usin gbcrypt library
func (user *User) ComparePassword(pass string) error {
	//bcrypt only works with slice of bytes data,hash the password received as parameter
	//and the pass returned by the DB
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
}
