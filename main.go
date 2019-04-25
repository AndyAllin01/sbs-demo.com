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
	r.HandleFunc("/signup", handlerFunc).Methods("POST")
//	r.HandleFunc("/signup", signupPost).Methods("POST")
	http.ListenAndServe(":8080", r)

}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SU")
	w.Header().Set("Content-Type", "text/html")
	err := tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		panic(err) //problem with template - no point continuing
	}
	//http.Redirect(w, r, r.Header.Get("Referer"), 302)

	//handlerFunc(w, r)
}

/*
//signupPost processes entered details and initiates automated email authorization
func signupPost(w http.ResponseWriter, r http.Request){

}*/