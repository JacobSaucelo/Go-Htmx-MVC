package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"webGoCrud/models"
)

type UsersType models.UsersType

func OpenSaveFile(fPath string) (UsersType, error) {
	saveFile, err := os.Open(fPath)
	if err != nil {
		fmt.Println("Error Opening save file")
		return UsersType{}, err
	}
	defer saveFile.Close()

	decoder := json.NewDecoder(saveFile)
	var save UsersType
	err = decoder.Decode(&save)
	if err != nil {
		fmt.Println("Error Opening save file")
		return UsersType{}, err
	}

	return save, nil
}
