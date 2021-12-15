package models

//User struct to create user table yn DB with automigration
type User struct {
	Id           uint
	FirstName    string
	LastName     string
	Email        string
	Password     string
	IsAmbassador bool
}
