package main

import (
	"html/template"
	"log"
	"net/http"
)

// Todo is the structure of the tasks
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData is que structure of the page
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
