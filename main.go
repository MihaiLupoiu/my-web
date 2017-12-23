package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	tmpl := template.Must(template.ParseFiles("./template/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ""
		tmpl.Execute(w, data)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
