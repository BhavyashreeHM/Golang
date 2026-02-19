package main

import (
	"fmt"
	hello "package/pack1"
	hey "package/pack1/subpack1"
	bye "package/pack2"
)

func main() {

	fmt.Println(hello.Hello("Hello Bhavya"))
	fmt.Println(hey.Hey("Hello Bhavyashree"))
	fmt.Println(bye.Bye("Hello Bhavyashree HM"))
}
