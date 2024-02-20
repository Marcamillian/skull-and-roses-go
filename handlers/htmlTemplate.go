package handlers

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func HandleTodoList(w http.ResponseWriter, r *http.Request) {

	listName := chi.URLParam(r, "listName")

	tmpl := template.Must(template.ParseFiles("templates/base.html"))

	// create the page data
	data := TodoPageData{
		PageTitle: listName,
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: false},
			{Title: "Task 3", Done: true},
		},
	}

	// render the data in the template
	tmpl.Execute(w, data)
}
