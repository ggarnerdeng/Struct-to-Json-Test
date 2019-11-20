package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Structs should represent the JSON string
type portfolioInfo struct {
	Name      string   `json:"name"`
	Title     string   `json:"title"`
	Email     string   `json:"email"`
	Phone     int      `json:"phone"`
	Aboutme   []string `json:"aboutme"`
	Education string   `json:"education"`
	Projects  []string `json:"projects"`
	Message   string   `json:"message"`
	Numero    int64    `json:"ref"` //Field with json:struct tag is stored with tag instead of var name.
}

func main() {

	//This encodes(marshals) struct to JSON
	john := portfolioInfo{
		Name:      "John Doe",
		Title:     "Software Programmer",
		Email:     "email.com",
		Phone:     1231234567,
		Aboutme:   []string{"Can code Golang", "Can implement HTML/CSS"},
		Education: "University of the West of Scotland",
		Projects:  []string{"Project 1", "Project 2", "Project 3"},
		Message:   "Hello World",
		Numero:    6669,
	}

	var jsonData []byte
	jsonData, err := json.MarshalIndent(john, "", "   ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
}
