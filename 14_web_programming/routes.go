package main

import "net/http"

func (app *application) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/about", app.About)
	mux.HandleFunc("/contact", app.Contact)
	return mux
}
