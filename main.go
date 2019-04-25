package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HB")
	err := tpl.ExecuteTemplate(w, "frontPage.gohtml", nil)
	if err != nil {
		panic(err) //problem with template - no point continuing
	}
}

func main() {
	fmt.Println("sbs-demo.com", time.Now()) // log message

	r := mux.NewRouter()
	r.HandleFunc("/", handlerFunc)
	r.HandleFunc("/signup", signup).Methods("GET")
	//r.HandleFunc("/signup", handlerFunc).Methods("POST")
	r.HandleFunc("/login", login).Methods("GET")
	r.HandleFunc("/login", loginPost).Methods("POST")
	r.HandleFunc("/signup", signupPost).Methods("POST")
	http.ListenAndServe(":8080", r)

}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SU")
	w.Header().Set("Content-Type", "text/html")
	err := tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		panic(err) //problem with template - no point continuing
	}
}

//signupPost processes entered details and initiates automated email authorization
func signupPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SP")
	http.Redirect(w, r, "/localhost:8080/", 200) // temporary redirect
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		panic(err) //problem with template - no point continuing
	}
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LP")
	http.Redirect(w, r, "localhost:8080/", 200) // temporary redirect
}
