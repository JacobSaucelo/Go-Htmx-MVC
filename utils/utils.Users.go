package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"webGoCrud/models"
)

type Users models.UsersType
type User models.UserType

var lastID int

func GetUsers() *Users {
	saveFile, err := os.Open("./data/data.json")
	if err != nil {
		generate := []byte(generateJson())

		folderError := os.MkdirAll("data", os.ModePerm)
		if folderError != nil {
			fmt.Println("cant create folder.")
		}

		fileErr := os.WriteFile("assets/store.json", generate, 0644)
		if fileErr != nil {
			fmt.Println("cant generate json file.")
		}
	}

	defer saveFile.Close()

	saveByteValue, _ := io.ReadAll(saveFile)

	var usersData Users

	json.Unmarshal(saveByteValue, &usersData)

	return &usersData
}

func generateJson() string {
	texts := `{"Users":[]}`
	// you can add default data

	return texts
}

func (u *Users) CreateUser(n string, a string) User {
	age, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Invalid age. Please enter a valid number.")
		// return
	}

	lastID++

	slug := User{
		Id:   lastID,
		Name: n,
		Age:  age,
	}

	u.Users = append(u.Users, models.UserType(slug))
	fmt.Println("user:", u.Users)

	// add it to db
	SaveSFile(UsersType(*u))

	return User{
		Id:   lastID,
		Name: n,
		Age:  age,
	}
}

func (u *Users) UpdateUser(id int, name string, age int) {
	var foundIndex int = -1
	// get json file

	for i, user := range u.Users {
		if user.Id == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("User not found.")
	}

	u.Users[foundIndex].Name = name
	u.Users[foundIndex].Age = age

	// SaveSFile(UsersType(*u))

	fmt.Println("Successfully updated user ", id)
}

// func (u *Users) DeleteUser(id string) Users {
func (u *Users) DeleteUser(id string) {
	var foundIndex int = -1

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error not a valid id.")
		return
	}

	for index, user := range u.Users {
		if user.Id == idInt {
			foundIndex = index
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("User not found.")
		return
	}

	u.Users[foundIndex] = u.Users[len(u.Users)-1]
	u.Users = u.Users[:len(u.Users)-1]

	SaveSFile(UsersType(*u))
	fmt.Println("Successfully removed user ", id)

}
