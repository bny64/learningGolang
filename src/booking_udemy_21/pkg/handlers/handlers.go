package handlers

import (
	"log"
	"net/http"

	"github.com/bny64/go-course/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested URL:", r.URL.Path)
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested URL:", r.URL.Path)
	render.RenderTemplate(w, "about.page.tmpl")

}
