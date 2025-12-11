package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("write.txt")
	if err != nil {
		fmt.Println("Error1 while opening file", err)
		return
	}
	defer func() {
		fmt.Println("Closing the file")
		file.Close()
	}()

	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error2 while reading the file")
		return
	}
	fmt.Println("File content:", string(data))
}
