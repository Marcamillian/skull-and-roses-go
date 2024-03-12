// Package main creates a webserver to run an implemenation
//
// This is a descirption of the package that we need

package main

import (
	"net/http"
	"time"

	"example.com/skull/handlers"
	"example.com/skull/wsServers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/net/websocket"
)

func main() {

	r := chi.NewRouter()

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
	r.Get("/todo/{listName}", handlers.HandleTodoList)

	// url parameters endpoint
	r.HandleFunc("/books/{title}/page/{page}", handlers.HandleBookRequest)

	// == endpoints for json encoding ==
	r.HandleFunc("/json/decode", handlers.HandleJsonDecode)
	r.HandleFunc("/json/encode", handlers.HandleJsonEncode)

	// golang.org chatserver
	golangServer := wsServers.NewWsGolangServerChat()
	r.Handle("/ws/golangChat", websocket.Handler(golangServer.HandleWSGolangChat))

	// set up static file serving
	fs := http.FileServer(http.Dir("static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// launch the server
	http.ListenAndServe(":3333", r)
}
