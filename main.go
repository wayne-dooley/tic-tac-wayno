package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Create a struct that holds information to be displayed in our HTML file
type TTW struct {
	Name string
}

// Go application entrypoint
func main() {

	//Instantiate a Welcome struct object and pass in some random infomation
	// We'll get the name of the user as a query parameter from the URL
	ttw := TTW{"Anomymous"}

	// Tell Go exactly where to find the HTML file
	// Ask Go to parse the HTML file
	templates := template.Must(template.ParseFiles("templates/ttw-template.html"))

	// Tell go to create a handle that looks into the 'static' directory
	// Go then uses the "/static/" as a URL that HTML can refer to when looking for CSS and other files
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// This method takes in the URL path "/" and a function that takes in a
	// response writer and a HTTP request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Takes the name fromthe URL query e.g. ?name=Martin,
		// sets welcome.name = Martin
		if name := r.FormValue("name"); name != "" {
			ttw.Name = name
		}
		// If errors, show an internal server error message
		// Also pass the welcome struct to the welcome-template.html file
		if err := templates.ExecuteTemplate(w, "ttw-template.html", ttw); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	// Start the web server, set the port to listen to 8080
	// Without a path, it assumes localhost
	// Print any errors from starting the webserver using fmt
	fmt.Println("Serving on port 8080..")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
