package main

import (
	"errors"
	"net/http"
)

func (app *application) Serve() error {
	if app.mux == nil {
		return errors.New("mux is not initialized!")
	}
	return http.ListenAndServe(":8080",app.mux)
}

func (app *application) mount(mux *http.ServeMux) {
	app.mux = mux

	app.mux.HandleFunc("/",app.Home)
}
