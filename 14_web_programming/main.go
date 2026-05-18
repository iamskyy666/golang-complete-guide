package main

import (
	"log"
	"net/http"
)

// ROUTING AND HTTP Handlers()

func main() {

	 mux:=http.NewServeMux()
	 mux.HandleFunc("/",Home)
	 mux.HandleFunc("/about",About)
	 mux.HandleFunc("/contact",Contact)

	 log.Println("Server listening on PORT: 8080 ✅")

	if err:= http.ListenAndServe(":8080",mux); err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}
}