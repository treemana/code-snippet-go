package main

import (
	"encoding/json"
	"fmt"
)

type omit bool

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ShotUser struct {
	Email string `json:"email"`
}

type NewUser struct {
	ShotUser
	Password *string `json:"password,omitempty"`
}

func main() {
	user := User{
		Email:    "email",
		Password: "password",
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	var nUser NewUser
	err = json.Unmarshal(bytes, &nUser.Email)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(nUser.Email)
	fmt.Println(nUser.Password)
}
