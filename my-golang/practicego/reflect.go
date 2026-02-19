package main

import (
	"fmt"
	"sync"
)

func Printvalue(wg *sync.WaitGroup, a interface{},ch chan string) {
	ch <- fmt.Sprintf("You entered value %d \n", a)
	<-ch
	Printvalue(wg,a,ch)
}
func PrintType(wg *sync.WaitGroup, a interface{},ch chan string) {
	wg.Done()
	fmt.Printf("of  type %T\n",a)
	<-ch
}
func hi() {
	var wg sync.WaitGroup
	email := "bhvyashreehm@gmail.com"
	// phone_number := 8197103471
	var ch chan string
	wg.Add(1)
	go Printvalue(&wg,email,ch)
	// Anyvalue(phone_number)
	wg.Wait()
	fmt.Println("Program execution completed")
}
