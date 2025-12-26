package ascii_art_web

import (
	"html/template"
	"net/http"
	"os"

	ascii_art "ascii_art_web/ascii-art/src"
)

// AsciiArtHandler handles the POST request to generate ASCII art.
// It expects form values "text" and "banner" and returns the rendered result in HTML.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST. If not, return a 400 Bad Request.
	templerr, _ := template.ParseFiles("error.html")
	if r.Method != http.MethodPost {
		data := map[string]string{
			"ErrorMessage": "BAD REQUEST",
		}
		w.WriteHeader(400)
		templerr.Execute(w, data)
		return
	}

	// Get form values from the request
	text := r.FormValue("text")

	if len(text) > 2000 {
		data := map[string]string{
			"ErrorMessage": "Max size 2000 charachter: 422",
		}
		w.WriteHeader(422)
		templerr.Execute(w, data)
		return
	}
	// The input text to convert to ASCII art
	banner := r.FormValue("banner") // The banner style to use

	// Validate that both fields are provided
	if text == "" || banner == "" {
		data := map[string]string{
			"ErrorMessage": "BAD REQUEST",
		}
		w.WriteHeader(400)
		templerr.Execute(w, data)
		return
	}

	// Read the banner file from the banners folder
	bannerData, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		// If the file doesn't exist or can't be read, return 404 Not Found
		data := map[string]string{
			"ErrorMessage": "BANNER NOT FOUND",
		}
		w.WriteHeader(404)
		templerr.Execute(w, data)
		return
	}

	// Convert the banner file content to a font usable by the ascii_art package
	font := ascii_art.Sep_Fonts(string(bannerData))
	// Generate ASCII art from the input text using the chosen font
	result := ascii_art.Chars_To_Art(font, text)
	if !IsSupportedText(result) {
		http.Error(w, "not validition text: 422", 422)
		return

	}

	// Parse the HTML template for rendering the result
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// If template is missing, return 404 Not Found
		http.Error(w, "Template not found: 404", http.StatusNotFound)
		return
	}
	// Prepare data to pass to the template
	data := map[string]string{
		"Result": result,
	}

	// Execute the template and write the result to the ResponseWriter
	if err := tmpl.Execute(w, data); err != nil {
		// If template execution fails, return 500 Internal Server Error
		http.Error(w, "Internal server error: 500", http.StatusInternalServerError)
		return
	}
}

func IsSupportedText(text string) bool {
	for _, v := range text {
		if v >= 32 && v <= 126 {
			return true
		}
	}
	return false
}
