package main

import (
	"html/template"
	"net/http"
	"path"
)

func (app *application) Render(w http.ResponseWriter, fileName string, data any) {
	fullPath:=path.Join(app.templateDir,fileName)
	tmpl,err:= template.ParseFiles(fullPath)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}