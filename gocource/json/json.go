package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Firstname string `json:"first_name"`
	Age       int    `json:"age, omitempty"`
	Salary    int    `json:"salary"`
}

func main() {
	p := Person{
		Firstname: "Bhavya",
		// Age: 24,
		Salary: 32000,
	}

	jsondata, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error while processing", err)
	}
	fmt.Println("JSon data\n", string(jsondata))
}
