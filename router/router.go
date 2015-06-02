package router

import (
	"net/http"
	"os"
	"text/template"
)

var templates = make(map[string]*template.Template)

func Router() {
	// Load Templates
	templates["index"] = template.Must(template.ParseFiles(getTemplatePath()+"/views/index.tpl", getTemplatePath()+"/views/base.tpl"))
	templates["register"] = template.Must(template.ParseFiles(getTemplatePath()+"/views/register.tpl", getTemplatePath()+"/views/base.tpl"))

	// Register Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/register", registerHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("query")

	data := struct {
		QueryParam string
	}{
		queryParam,
	}
	err := templates["index"].ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["register"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTemplatePath() string {
	return os.Getenv("GOPATH") + "/src/github.com/kusold/talks-i-want-to-hear"
}
