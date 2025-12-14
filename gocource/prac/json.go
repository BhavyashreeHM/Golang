package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	firstname string `json:"first_name"`
	age       int    `json:age`
}

func main() {
	p := person{
		firstname: "Bhavya",
		age:       25,
	}
	jsondata, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(string(jsondata))

}
