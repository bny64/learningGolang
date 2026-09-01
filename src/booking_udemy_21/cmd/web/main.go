package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bny64/go-course/pkg/config"
	"github.com/bny64/go-course/pkg/handlers"
	"github.com/bny64/go-course/pkg/render"
)

const portNumber = ":8080"

// main is the application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
