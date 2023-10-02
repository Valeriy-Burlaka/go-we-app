package main

import (
	// "encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Global variable to hold compiled templates
var templates *template.Template

type Address struct {
	Country string `json:country`
	City string `json:city`
	Street string `json:street`
	House string `json:house`
}

func addressHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// This is one way of doing things: parsing a template at runtime, when the handler is invoked:
		// 
		// The function `template.Must` will panic if the template.ParseFiles function returns an error.
		// // tmpl := template.Must(template.ParseFiles("html/index.html"))
		// 
		// The 2nd argument to `tmpl.Execute(w, nil)` is the data that the template will use when it's executed.
		// We can pass in any dynamic content that we want to render in the template.
		// For example, if the template has placeholders like `{{.Country}}` for displaying a country,
		// we could pass in an object with a Country field:
		// ```
		// 	data := Address{Country: "Poland", City: "Krakow"}
		// 	tmpl.Execute(w, data)
		// ```
		// 
		// // tmpl.Execute(w, nil)

		// The 2nd way is to preliminary parse the template inside the `init()` function,
		// so we can catch parse template errors before starting the service.
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func init() {
	var err error
	templates, err = template.ParseFiles("html/index.html", "html/create-address-form.html")
	if err != nil {
		log.Fatalf("Error parsing HTML templates: %v", err)
	}
}

func main() {
	port := ":8101"

	http.HandleFunc("/address", addressHandler)
	// The 2nd argument allows to specify other HTTP handler.
	// By default, it's equal to `DefaultServerMux`, which is a global variable, a singleton instance of `http.ServeMux`
	fmt.Printf("Starting HTTP server on port %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Server failed to start:", err, err.Error())
	}
}
