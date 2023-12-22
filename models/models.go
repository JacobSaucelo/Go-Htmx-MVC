package models

type UserType struct {
	Id   int    `json:id`
	Name string `json:name`
	Age  int    `json:age`
}

type UsersType struct {
	Users []UserType `json:users`
}
