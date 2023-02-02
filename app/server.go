package main

import (
	"fmt"
	// "log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/posts", postsHandler)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ハロー・ワールド")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "users")
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "posts")
}
