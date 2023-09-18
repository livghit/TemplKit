package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"

	button "github.com/livghit/templkit/components/Actions/button"
	"github.com/livghit/templkit/pages"
)

func main() {

	if err := run(); err != nil {
		log.Fatalf("Error running the server : %s", err)
	}

}

func run() error {
	o := button.ButtonOptions{
		Label: "hello",
		Class: "btn btn-primary",
	}

	//mux router

	router := mux.NewRouter()
	index := pages.Index(o)
	notFound := pages.NotFound()

	router.Handle("/", templ.Handler(index))
	router.NotFoundHandler = templ.Handler(notFound)

	log.Info("Server started at 8080")
	return http.ListenAndServe(":8080", router)
}
