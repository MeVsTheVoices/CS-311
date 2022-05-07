package main

import (
	"fmt"
	"log"
	"net/http"
)

var emails map[string]string

//convenience method for extracting fields
func getFieldFromURL(w http.ResponseWriter, r *http.Request, field string) (string, bool) {
	name := r.URL.Query().Get(field)
	//if the field wasn't specified, send back and error describing
	if len(name) < 1 {
		http.Error(w, field+" wasn't specified", http.StatusNotFound)
		return "", false
	}
	return name, true
}

//associated with /lookup
func lookupHandler(w http.ResponseWriter, r *http.Request) {
	name, foundFieldValue := getFieldFromURL(w, r, "name")
	if !foundFieldValue {
		return
	}

	//check if the email has an entry and serve if it does
	email, nameExists := emails[name]
	if nameExists {
		fmt.Fprintf(w, email)
	} else {
		fmt.Fprintf(w, name+" doesn't have an email associated")
	}
}

//associated with /add
func addHandler(w http.ResponseWriter, r *http.Request) {
	name, foundNameValue := getFieldFromURL(w, r, "name")
	email, foundEmailValue := getFieldFromURL(w, r, "email")

	//if we didn't find a value for either name or email, bail out
	if !foundNameValue || !foundEmailValue {
		return
	}

	//add and entry for name, overrite if existing
	emails[name] = email

	//notify user of changes
	fmt.Fprintf(w, "added "+name+" as "+email)
}

func main() {
	//create a map for emails and associate handler functions
	emails = make(map[string]string)
	http.HandleFunc("/lookup", lookupHandler)
	http.HandleFunc("/add", addHandler)

	//sit back and wait for requests
	log.Fatal(http.ListenAndServe(":8080", nil))
}
