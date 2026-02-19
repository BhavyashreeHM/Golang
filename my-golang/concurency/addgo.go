package main

import (
	"fmt"
)

func main() {
	var a, b int
	c := make(chan int)
	fmt.Print("Enter two numbers: ")
	fmt.Scanf("%d\n %d\n", &a, &b)
	
	go add(c, a, b)
	fmt.Println("Sum", <-c)
	fmt.Println(a, "+", b, "=", Add(a, b))
}

func add(c chan int, a int, b int) {
	c <- a + b
}

func Add(a, b int) int {
	return a + b
}
