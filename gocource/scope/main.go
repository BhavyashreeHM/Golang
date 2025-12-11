package main

import "fmt"

var fname string = "Bhavya"

func main() {
	middlename := "shree"

	lname := pritname()

	fmt.Println(fname+middlename + " "+lname)

	fname = "Kavya"
	fmt.Println(fname + " " + middlename + " " + lname)
}

func pritname() string {
	Lastname := "HM"
	return Lastname
}
