package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing template:", err)
		return
	}
}
