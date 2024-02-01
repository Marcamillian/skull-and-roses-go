package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {

	r := chi.NewRouter()
	tmpl := template.Must(template.ParseFiles("templates/base.html"))

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped
	r.Use(middleware.Timeout(60 * time.Second))

	// root page handler
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	// template endpoint
	r.Get("/todo/{listName}", func(w http.ResponseWriter, r *http.Request) {

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

		// render the data in the template
		tmpl.Execute(w, data)
	})

	// url parameters endpoint
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {

		title := chi.URLParam(r, "title")
		page := chi.URLParam(r, "page")

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

	})

	// set up static file serving
	fs := http.FileServer(http.Dir("static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// launch the server
	http.ListenAndServe(":3333", r)
}
