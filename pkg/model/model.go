package model

import (
	"encoding/json"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Rows    []User `json:"rows"`
}

type Substring struct {
	Substring string `json:"substring"`
}

type EmailCheck struct {
	Email string `json:"Email"`
}

type IinCheck struct {
	Iin string `json:"Iin"`
}
type IinCheckN struct {
	IinN json.Number `json:"Iin"`
}
