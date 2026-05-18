package main

import (
	"fmt"
	"log"
	"net/http"
)

var htmlTmp =`
		<!DOCTYPE html>
		<html>
		<head>
		<title>%s</title>
		</head>
		<body>
		%s
		</body>
		</html>
		`
// dummy handlers - ROUTING AND HTTP Handlers()

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s %s",r.Method,r.URL.Path)
	if r.Method == http.MethodPost{
		// ... some form-processing here
	}
	app.Render(w,"index.html",nil)
}

func (app *application) About(w http.ResponseWriter, r *http.Request) {
	aboutContent:= `<h2>About</h2>
	<div>We are a small shop, doing great things!</div>
	`
	aboutContent = fmt.Sprintf(htmlTmp,"About us",aboutContent)
	_,err:=w.Write([]byte(aboutContent))
	if err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}
}

func (app *application) Contact(w http.ResponseWriter, r *http.Request) {
	contactContent:= `
	<h2>Contact</h2>
	<div>send as an email on blablabla@test.com</div>
	`
	contactContent = fmt.Sprintf(htmlTmp,"Contact us",contactContent)
	_,err:=w.Write([]byte(contactContent))
	if err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}
}