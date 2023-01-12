package models

type Product struct {
	Id   int    `json: "id"`
	NAME string `json: "name"`
	TYPE string `json: "type"`
}
