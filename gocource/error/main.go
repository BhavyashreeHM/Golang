package main

import (
	"fmt"
	"strings"
)

func main() {
	sr := "He"
	fmt.Println(strings.Contains(sr, "h"))
	name := []string{"A","B","C"}
	a:= strings.Join(name, "D")
	fmt.Println(a)

}
