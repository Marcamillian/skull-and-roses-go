package handlers

import (
	"html/template"
	"net/http"

	"example.com/skull/templates"

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

	var importedTemplate *template.Template = templates.GetUsableTemplate("templateFiles/base.go.html")

	listName := chi.URLParam(r, "listName")

	// create the page data
	data := TodoPageData{
		PageTitle: listName,
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: false},
			{Title: "Task 3", Done: true},
		},
	}

	importedTemplate.Execute(w, data)
}
