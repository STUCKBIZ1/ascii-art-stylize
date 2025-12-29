package ascii_art_web

import (
	"html/template"
	"net/http"
)

// HomeHandler handles the GET request to render the home page.
// It simply serves the index.html template with an empty Result field.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorStyleHandler(w, "THAT PAGE NOT FOUND: 404", 404)
		return
	}
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		// If not GET, return 400 Bad Request
		ErrorStyleHandler(w, "BAD REQUEST: 400", 400)
		return
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// If template is missing, return 404 Not Found
		ErrorStyleHandler(w, "TEMPLATE NOT FOUND: 404", 404)
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
		ErrorStyleHandler(w, "INTERNAL SERVER ERROR: 500", 500)
		return
	}
}
