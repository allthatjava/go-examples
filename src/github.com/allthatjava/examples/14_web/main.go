package main

import (
	"fmt"
	"net/http"
)



func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello Golang World")
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>About</h1")
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}
