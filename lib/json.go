package lib

import (
	"encoding/json"
	"fmt"
)

/*
Marshal => Converting go object into json string
Unmarshall => converting json string to go object
*/

//JSON func
func JSON() {
	jsonRead()
}

//SensorReading struct
type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func jsonRead() {
	jsonString := `{"name":"Angel Dhakal", "capacity":40, "time":"09:15"}`
	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}

//Book struct
type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

//Author struct
type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func book() {
	author := Author{
		Name:      "Angel Dhakal",
		Age:       19,
		Developer: true,
	}
	book := Book{
		Title:  "Learning Cuncurrency",
		Author: author,
	}
	fmt.Printf("%+v\n", book)

	byteSlice, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteSlice))
}
