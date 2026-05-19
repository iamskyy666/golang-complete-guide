package main

import (
	"net/http"
)

func (app *application) Routes() http.Handler {
	mux := http.NewServeMux()

	// staticDir:=filepath.Join(".","14_web_programming","public")
	mux.Handle("/public/",http.StripPrefix("/public/",http.FileServer(http.Dir(app.publicPath))))

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/about", app.About)
	mux.HandleFunc("/contact", app.Contact)

	return mux
}
