package main

import (
	_ "embed"
	"fmt"
)

//go:embed Output.txt
var content string

func main() {
	fmt.Println("Embeded content\n", content)
}
