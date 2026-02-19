package main

import "fmt"

func main() {

	var num int = 100
	var str string = "Golang"
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(num, str, arr)

	var number *int
	var name *string

	number = &num
	name = &str
	fmt.Println("Variable ADDRESS:Here the number is stored",number)
	fmt.Println("Here the name is stored",name)

	fmt.Println("pOINTER ADDRESS:Here the pointer number is stored",&number)
	fmt.Println("Here the pointer name is stored",&name)

	fmt.Println(*number)
	fmt.Println(*name)

	

	var a= 100

	var b=&a
	fmt.Println(a,b)

	x := "Apple"

	y :=&x

	fmt.Println(x,y,*y)

	*y ="Ball"
	fmt.Println(*y)

	bl := true
	var n *bool = &bl
	fmt.Println(*n)
	*n = false
	fmt.Println(*n)

	var username *string
	fmt.Println(username)
	fmt.Println(*username)
}
