package main

import (
	"fmt"
	intslice "helper/slice/intSlice"
	str "helper/string"
)

func main() {
	fmt.Println(intslice.SortInt([]int{14, 5, 19, 10}))
	fmt.Println(str.Upper("golinuxcloud"))
}
