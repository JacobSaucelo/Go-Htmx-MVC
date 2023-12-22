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

func GetUsers() models.UsersType {
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

	return models.UsersType(usersData)
}

func generateJson() string {
	texts := `{"Users":[]}`
	// you can add default data

	return texts
}

func (u *Users) CreateUser(n string, a string) {
	age, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Invalid age. Please enter a valid number.")
		return
	}

	lastID++

	slug := User{
		Id:   lastID,
		Name: n,
		Age:  age,
	}

	u.Users = append(u.Users, models.UserType(slug))

	// add it to db

	fmt.Println("User created successfully.")
}
