package openframe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	Id string
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Save(filePath string) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err.Error())
	}

	return ioutil.WriteFile(filePath, bytes, 0644)
}

func LoadUser(filePath string) (User, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return User{}, err
	}

	var user User
	json.Unmarshal(file, &user)

	return user, nil
}
