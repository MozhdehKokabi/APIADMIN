package models

type User struct {
	Password string
	UserName string
	Email    string
	Phone    int
	Address  string
	Role     string
	Id       uint
}

type Admin struct {
	Password string
	UserName string
	Email    string
	Role     string
	Id       uint
}

type ReqUser struct {
	Password string `json:"password"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Address  string `json:"address"`
}

type ReqAdmin struct {
	Password string `json:"password"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
}

type ReqWebsite struct {
	Domain  string `json:"domain"`
	Name    string `json: "name"`
	Address string `json:"address"`
}
