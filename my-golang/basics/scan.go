package main

import "fmt"

func main() {
	fmt.Println("What is your name?")
	
	var name string
	// fmt.Scanln(&name)
	fmt.Scan(&name)

	fmt.Printf("Hello %s, welcome to Go programming! \n", name)

}
