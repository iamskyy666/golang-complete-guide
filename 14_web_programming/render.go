package main

import (
	"net/http"
)

func (app *application) Render(w http.ResponseWriter, fileName string, data any) {
	if app.tp==nil{
		http.Error(w,"template rendering engine FAILURE", http.StatusInternalServerError)
		return
	}
	app.tp.Render(w,fileName,data)
}