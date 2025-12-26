package ascii_art_web

import (
	"html/template"
	"net/http"
)

// HomeHandler handles the GET request to render the home page.
// It simply serves the index.html template with an empty Result field.
func HomeHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.Error(w, "not found: 404", 404)
		return
	}
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		// If not GET, return 400 Bad Request
		http.Error(w, "Bad Request: 400", http.StatusBadRequest)
		return
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// If template is missing, return 404 Not Found
		http.Error(w, "Template not found: 404", http.StatusNotFound)
		return
	}

	// Prepare empty data to pass to the template
	data := map[string]string{
		"Result": "", // No ASCII art yet
	}

	// Execute the template and write it to the ResponseWriter
	// Good practice: handle possible execution error
	if err := tmpl.Execute(w, data); err != nil {
		// If template execution fails, return 500 Internal Server Error
		http.Error(w, "Internal server error: 500", http.StatusInternalServerError)
		return
	}
}
