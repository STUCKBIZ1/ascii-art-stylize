package ascii_art_web

import (
	"html/template"
	"net/http"
)

func ErrorStyleHandler(w http.ResponseWriter, errorstring string, codestatut int) {
	tmpl, _ := template.ParseFiles("templates/error.html")
	data := map[string]string{
		"ErrorMessage": errorstring,
	}
	w.WriteHeader(codestatut)
	tmpl.Execute(w, data)
}
