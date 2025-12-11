package file

import (
	"fmt"
	"os"
)

func main() {
	
	file, err := os.Create("Output.txt")
	if err != nil {
		fmt.Println("Error while processesing",err)
	}
	defer file.Close()

	data := []byte("Hello world\n\n")
	_, err = file.Write(data)
		if err != nil {
		fmt.Println("Error while processesing",err)
	}

	fmt.Println("Data successfully written to file")
}
