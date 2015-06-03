package router

import (
	"database/sql"
	"net/http"
	"os"
	"text/template"

	_ "github.com/lib/pq"
)

var templates = make(map[string]*template.Template)

func Router() {
	// Load Templates
	templates["index"] = template.Must(template.ParseFiles(getTemplatePath()+"/views/index.tpl", getTemplatePath()+"/views/base.tpl"))
	templates["register"] = template.Must(template.ParseFiles(getTemplatePath()+"/views/register.tpl", getTemplatePath()+"/views/base.tpl"))
	templates["register-confirm"] = template.Must(template.ParseFiles(getTemplatePath()+"/views/register-confirm.tpl", getTemplatePath()+"/views/base.tpl"))
	templates["login"] = template.Must(template.ParseFiles(getTemplatePath()+"/views/login.tpl", getTemplatePath()+"/views/base.tpl"))
	templates["login-confirm"] = template.Must(template.ParseFiles(getTemplatePath()+"/views/login-confirm.tpl", getTemplatePath()+"/views/base.tpl"))

	// Register Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
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
	switch r.Method {
	case "GET":
		registerGetHandler(w, r)
	case "POST":
		registerPostHandler(w, r)
	}
}
func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["register"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	var user = struct {
		Email    string
		Password string
	}{
		r.FormValue("email"),
		r.FormValue("password"),
	}

	db, err := sql.Open("postgres", "user=postgres dbname=golang sslmode=disable")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.Ping()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users(email, password) VALUES($1,$2)", user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates["register-confirm"].ExecuteTemplate(w, "base", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		loginGetHandler(w, r)
	case "POST":
		loginPostHandler(w, r)
	}
}
func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["login"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	var user = struct {
		Email    string
		Password string
	}{
		r.FormValue("email"),
		r.FormValue("password"),
	}

	db, err := sql.Open("postgres", "user=postgres dbname=golang sslmode=disable")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.Ping()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var userResult = struct {
		ID       int
		Email    string
		Password string
	}{}
	err = db.QueryRow("SELECT * FROM users WHERE email = $1 AND password = $2", user.Email, user.Password).Scan(&userResult.ID, &userResult.Email, &userResult.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates["login-confirm"].ExecuteTemplate(w, "base", userResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTemplatePath() string {
	return os.Getenv("GOPATH") + "/src/github.com/kusold/talks-i-want-to-hear"
}
