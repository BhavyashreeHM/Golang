package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Project struct {
	ID   int    `json:"id"`
	Tech string `json:"tech"`
}

var (
	nextId   = 1
	projects = []Project{{
		ID:   nextId,
		Tech: "Golang",
	}}
)

// go mod init myapi
func main() {
	nextId++
	mux := http.NewServeMux()

	mux.HandleFunc("/crud/", Handler)
	mux.HandleFunc("/crud/create", CHandler)
	mux.HandleFunc("/crud/read", RHandler)
	mux.HandleFunc("/crud/update/{id}", Updatehandler)
	mux.HandleFunc("/crud/delete/{id}", Dhandler)

	log.Println("Server Up and Running on PORT 3000")
	http.ListenAndServe(":3000", mux)
	// http.ListenAndServe(":8000", mux)
}
func CHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var project []Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range project {
		project[i].ID = nextId
		nextId++
		projects = append(projects, project[i])
	}
	fmt.Println(nextId)
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Go CRUD API")
}

func RHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)

}

func Updatehandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// idS := r.PathValue("id")
	// id, err := strconv.Atoi(idS)
	idStr := strings.TrimPrefix(r.URL.Path, "/crud/update/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprint(w, "Id is not found")
		// w.WriteHeader(http.StatusNotFound)
		// json.NewEncoder(w).Encode(map[string]string{
		// 	"error": "Invalid user ID",
		// })

		return
	}
	for i, project := range projects {
		if project.ID == id {
			err = json.NewDecoder(r.Body).Decode(&projects[i])
			projects[i].ID = id
			json.NewEncoder(w).Encode(projects[i])
			return
		}
	}
}

func Dhandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/crud/delete/")
    id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprint(w, "Id is not found")
		// w.WriteHeader(http.StatusNotFound)
		// json.NewEncoder(w).Encode(map[string]string{
		// 	"error": "Invalid user ID",
		// })

		return
	}

    for i, user := range projects {
        if user.ID == id {
            projects = append(projects[:i], projects[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }

}
}

