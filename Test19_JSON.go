package main

import (
	"encoding/json"
	"fmt"
)

type newUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	var u = newUser{1, "name", "email"}
	res, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
