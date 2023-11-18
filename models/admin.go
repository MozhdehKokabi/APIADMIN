package models

type Admin struct {
	Password string
	UserName string
	Email    string
	Role     string
	Id       uint
}


type ReqAdmin struct {
	Password string `json:"password"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
}