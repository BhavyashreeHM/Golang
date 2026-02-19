package main

import (
	greeting "golinuxclod/greeting"
	"log"
	"os"
)

func main() {
	value := os.Args[1]
	res, err := greeting.Greeting(value)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	log.Println(res)
}
