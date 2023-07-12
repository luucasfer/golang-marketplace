package main

import (
	"html/template"
	"net/http"
)

var readTemplates = template.Must(template.ParseGlob("templates/*.html")) //le todos templates que tem .html

func main() {
	http.HandleFunc("/", index)   //toda vez que chamar a barra, executo a func index
	http.ListenAndServe(":8000", nil)
}


func index(w http.ResponseWriter, r *http.Request){  //(w writer, r responser)
	readTemplates.ExecuteTemplate(w, "index", nil)
}