package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// temp db
var Users = map[string]string{}
