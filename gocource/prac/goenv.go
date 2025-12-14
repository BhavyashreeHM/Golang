package env

import (
	"fmt"
	"os"
)

func main() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env variable", user)
	fmt.Println("Home env variable", home)
}
