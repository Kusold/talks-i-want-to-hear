package router

import (
	"net/http"
	"os"
	"text/template"
)

func Router() {
	http.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("query")
	tpl, err := template.ParseFiles(getTemplatePath() + "/views/base.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		QueryParam string
	}{
		queryParam,
	}
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTemplatePath() string {
	return os.Getenv("GOPATH") + "/src/github.com/Kusold/talks-i-want-to-hear"
}
