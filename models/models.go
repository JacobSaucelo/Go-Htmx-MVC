package models

type UserType struct {
	Id   int
	Name string
	Age  int
}

type UsersType struct {
	Users []UserType
}
