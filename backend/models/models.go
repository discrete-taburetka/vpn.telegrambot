package models

// import "gorm.io/gorm"

//Tg_users model
type Tg_users struct {
	Tg_user_name  	string `json:"tg_user_name"`
	User_name     	string `json:"user_name"`
    User_password 	string `json:"user_password"`
    User_email    	string `json:"user_email"` 
}