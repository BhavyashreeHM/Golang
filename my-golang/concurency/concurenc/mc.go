package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var username, password string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your username ")
	// fmt.Scanf("%s", &username)
	username, _ = reader.ReadString('\n')
	
	Login(username)

	fmt.Println("Please enter your password ")
	fmt.Scanf("%s", &password)
	go Login(username)
	time.Sleep(1 * time.Second)

	fmt.Println("Please follow me on Instagram")
}
func Login(username string) {
	var oddSum string
	fmt.Printf("%s Login succesfull\n",username)
	fmt.Println("Solve this puzzle")
	fmt.Println("What is the sum of two odd number")
	fmt.Scan("%s",&oddSum)
	fmt.Println("thanks for participating")
}
