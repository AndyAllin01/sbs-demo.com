package main

/*
- Very basic code model for sbs tech task.
- Pseudo code and functions to be filled out and tidied.
- Refactoring into packages and some functional reusability is required.
- Have made this partially runnable to show /index homepage with register
    and login routing.

- I have not used Quasar or Echo, so this initial draft is using Go with
	gorilla mux, bootstrap HTML/CSS and postgresql.

- The requirements suggest mutliple users can have access to the same list,
	so there will be a need to query the user to find which list they want to
	access (together with some sort of list ownership and permissions?). I'd
	not see a to do list as an application requiring concurrent access so while
	goroutines amd channels may not be necessary, mutex locking to ensure there
	is no inadvertent database manipulation would be sensible (use mutex).

- I am aware of JWT but never used them (will be added to my own "to do" list!)

- This source in github https://github.com/AndyAllin01/sbs-demo.com
	I use github minimally for personal repo.

- I have not deployed any scalable apps, but understand that Docker containers managed
	by Kubernetes seems to be the method of choice at the moment. I'm very interested
	in learning more about this and have read and viewed some high level summaries.
*/

import (
	"fmt"
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq" //use postgresql to store/validate users and save lists
)

type list struct {
	listName     string     // name of this list
	listUsers    []string   // slice of users allowed access to this list
	listContents []listToDo // list details
}

//toDo is top-level list linking user (ID) to a list
type toDo struct {
	ID     string //unique email address
	events []listToDo
}

//listToDo is a possible sub-list
type listToDo struct {
	thingToDo string
	emement   []string
}

var tpl *template.Template

var mtx sync.Mutex // can multiple users access the same list?

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	/*	psqlDetail := "postgres://user:password@database_location?sslmode=disable"

		db, err := sql.Open("postgres", psqlDetail)
		if err != nil {
			panic(err)
		}*/
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
	r.HandleFunc("/signup", signupPost).Methods("POST")
	r.HandleFunc("/login", login).Methods("GET")
	r.HandleFunc("/login", loginPost).Methods("POST")
	http.ListenAndServe(":8080", r)
}

/*
FUNCTIONS BELOW HERE TO BE REFACTORED INTO SEPARATE PACKAGE(S)
==============================================================
*/

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
	http.Redirect(w, r, "/", http.StatusSeeOther) // temporary redirect
	//check for duplicate email
	//if unique, add to database user table
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
	http.Redirect(w, r, "/", http.StatusSeeOther) // temporary redirect
	//check if username and password is valid
}

//getList receives a validated user id (email address string)
//and returns that user's list
func getList(id string) (*listToDo, error) {
	return nil, nil
}

//add an element to the list
func add(li *listToDo) (*listToDo, error) {
	return li, nil
}

//delete an element from the list
func delete(li *listToDo) (*listToDo, error) {
	return li, nil
}
