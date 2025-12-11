package write

import (
	"bufio"
	"fmt"
	"go/scanner"
	"os"
)

func main() {
	file, err := os.Create("write.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error1 while processing", err)
		return
	}

	// //read the content of opened file
	// _, err = file.WriteString("Hello golang \n")
	// if err != nil {
	// 	fmt.Println("Error2 while processing", err)
	// 	return
	// }
	// fmt.Println("Data succesfully written")


	// using bufio
	scanner := bufio.NewScanner()
	for scanner.Scan(){
		line := scanner.Text()
		fmt.Println("Line", line)
	}
	err = scanner.Err() 
	if err != nil{
		fmt.Println("Error while reading")
	}

}
