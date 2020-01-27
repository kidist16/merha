package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("delivery/web/templates/*"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/index", indexHandler)

	http.ListenAndServe(":8080", mux)
}
