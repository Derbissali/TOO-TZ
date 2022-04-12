package model

import (
	"encoding/json"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"first_name"`
	Surname string `json:"last_name"`
}
type UpdateU struct {
	Name    string `json:"first_name"`
	Surname string `json:"last_name"`
}
type Substring struct {
	Substring string `json:"substring"`
}
type SubstringN struct {
	SubstringN json.Number `json:"substring"`
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

type Crc64 struct {
	ID   string `json:"id"`
	Hash string
}
