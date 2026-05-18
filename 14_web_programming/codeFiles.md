```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// 1️⃣ INTRO. TO ROUTING
// routing - mux
// routing - handlers (controllers in NodeJs/ExpressJs)
// GET / - homePage
// POST /users - CreateUser()

type DefaultHandler struct{}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	data:=map[string] interface{}{
		"user":"Skyy Banerjee",
		"age":30,
		"height":6,
		"location":"Kolkata",
	}

	bSlice,err:=json.Marshal(data)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

		// w.Header().Set("Content-Type","text/html; charset=utf-8")
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		// w.Write([]byte("<h1>Hello World!</h1>"))
		 w.Write(bSlice)
}

func main() {

	// mux:=http.NewServeMux()
	// mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
	// 	// w.Write([]byte("✅ Backend-Server running!"))
	// 	w.Header().Set("Content-Type","text/html; charset=utf-8")
	// 	htmlContent:= `
	// 	<!DOCTYPE html>
	// 	<html>
	// 	<head>
	// 	<title>About Us</title>
	// 	</head>
	// 	<body>
	// 	<h1>About this simple server</h1>
	// 	<p>This server is built using Golang's 'net/http' pkg.</p>
	// 	</body>
	// 	</html>
	// 	`
	// 	// fmt.Fprint(w,htmlContent)
	// 	_,err:=w.Write([]byte(htmlContent))
	// 	if err != nil{
	// 		fmt.Println("ERROR:",err.Error())
	// 		return
	// 	}
	// })

	mux:=&DefaultHandler{} // better way

	if err:= http.ListenAndServe(":8080",mux); err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}
}


// http://localhost:8080/
/*
{
"age": 30,
"height": 6,
"location": "Kolkata",
"user": "Skyy Banerjee"
}
*/
```