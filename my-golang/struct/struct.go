package main

import (
	"encoding/json"
	"fmt"
)

// Server stuct any type
type Server struct {
	Id   int    `json:"id"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

var (
	srv    = make(map[int]Server) // map of servers key:value example 1:Server, 2:Server
	nextID = 1
)

func init() {
	srv[nextID] = Server{

		Id:   nextID,
		Host: "localhost",
		Port: 8080,
	}
	nextID++
	srv[nextID] = Server{
		Id:   nextID,
		Host: "example.com",
		Port: 9090,
	}
	nextID++
}

func main() {
	serverlist := make([]Server, 0, len(srv)) // slice of servers
	for _, s := range srv {
		serverlist = append(serverlist, s)
	}
	fmt.Println(serverlist)
	jsonData, err := json.MarshalIndent(serverlist, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

}
