package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"webGoCrud/models"
)

type UsersType models.UsersType

var folderName string = "data"

func checkDirExists() {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, os.ModePerm)
		if err != nil {
			fmt.Println("error creating data folder", err)
			return
		}
	}
}

func ReadSaveFile(fPath string) (UsersType, error) {
	checkDirExists()
	saveFile, err := os.Open(fPath)
	if err != nil {
		fmt.Println("Error Opening save file.")
		return UsersType{}, err
	}
	defer saveFile.Close()

	decoder := json.NewDecoder(saveFile)
	var save UsersType
	err = decoder.Decode(&save)
	if err != nil {
		fmt.Println("Error reading save file.")
		return UsersType{}, err
	}

	return save, nil
}

func SaveSFile(fPath string, data UsersType) error {
	checkDirExists()
	saveFile, err := os.Create(fPath)
	if err != nil {
		fmt.Println("Error creating save file", err)
		return err
	}

	defer saveFile.Close()

	encoder := json.NewEncoder(saveFile)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding save file", err)
		return err
	}

	return nil
}
