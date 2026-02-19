package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the reating of our pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating our pizza:", input)

	fmt.Printf("The type of input is %T \n",input)

	numberOfRating, err := strconv.ParseFloat(strings.TrimSpace(input) ,64)

	if err!= nil{
		fmt.Println("Error")
	}else{
		fmt.Printf("The rating of pizza is %.2f \n",numberOfRating+1)
	}
	// input1,_ := reader.Read()
	// fmt.Println(string(input1))
}
