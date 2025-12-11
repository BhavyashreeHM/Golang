package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Error while processing", err)
	}
	defer file.Close()

	data := []byte("This is my first file")
	_, err= file.Write(data)
	if err!= nil{
		fmt.Println("Error while procesing",err)
	}
	fmt.Println("Sucesfully wrote data")
}
