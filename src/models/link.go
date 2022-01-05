package models

//struct to create a many to many relationship between user and product tables
type Link struct {
	Model
	Code     string    `json:"code"`
	UserId   uint      `json:"user_id"`
	User     User      `json:"user" gorm:"foreignKey:UserId"`
	Products []Product `json:"products" gorm:"many2many:link_products"`
}
