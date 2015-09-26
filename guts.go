
package main

import (
	"net/http"
	"html/template"
)

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("results.html")
	t.Execute(w, r) // I put r here, because it needs 2 arguments, but I have no idea why this is ok
}



func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("main.html")
	t.Execute(w, r) // I put r here, because it needs 2 arguments, but I have no idea why this is ok
}


func main() {
	http.HandleFunc("/results", resultsHandler)
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe("localhost:3000", nil)
}