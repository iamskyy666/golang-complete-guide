package main

import (
	"html/template"
	"net/http"
	"path"
	"path/filepath"
	"sync"
)

// Template-rendering with caching in Golang 🧠

type TmpltRenderer struct {
	cache map[string]*template.Template
	mutex sync.RWMutex
	dev bool
	tmpltDir string
}

func NewTemplateRenderer(tmpltDir string, isDev bool)*TmpltRenderer{
	return &TmpltRenderer{
		tmpltDir: tmpltDir,
		cache: make(map[string]*template.Template),
		dev: isDev,
		// For mutex, we don't need anything, it can be used by-default
	}
}

func (t *TmpltRenderer) Render(w http.ResponseWriter, templateName string, data any){
	tmplt,err:=t.getTemplate(templateName)

	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmplt.ExecuteTemplate(w,"base.html" ,data)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TmpltRenderer) getTemplate(tmpltName string)(*template.Template, error){
	if !t.dev {

	t.mutex.RLock()

	if tmpl, ok := t.cache[tmpltName]; ok {
		t.mutex.RUnlock()
		return tmpl, nil
	}

	t.mutex.RUnlock()
}


	// if tmplt not in cahce, we parse it
	tmpl,err:=t.parseTmplt(tmpltName)
	if err!=nil{
		return nil,err
	}

	if !t.dev{
		t.mutex.Lock()
		t.cache[tmpltName]=tmpl
		t.mutex.Unlock()
	}

	return tmpl,nil
}


func (t *TmpltRenderer) parseTmplt(tmpltName string)(*template.Template, error){
	tmpltPath:= path.Join(t.tmpltDir,tmpltName)

	files:=[]string{tmpltPath}

	layoutPath:=path.Join(t.tmpltDir,"layouts/*.html")
	layouts,err:=filepath.Glob(layoutPath)

	if err == nil{
		files = append(files, layouts...)
	}

	partialPath:=path.Join(t.tmpltDir,"partials/*.html")
	partials,err:=filepath.Glob(partialPath)

	if err == nil{
		files = append(files, partials...)
	}

	tmpl,err:=template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	// finally..
	return tmpl,nil
}