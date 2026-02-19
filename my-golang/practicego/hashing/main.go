package main

import (
	"fmt"
)

func main() {
	var password string
	fmt.Println("Please enter the password :")
	fmt.Scanln(&password)

	hash, _ := Hash(password)

	fmt.Println("Password", password)
	fmt.Println("Hash", hash)

	match := Verify(password, hash)
	fmt.Println("Match", match)

}
