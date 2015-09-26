
package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("main.html")
	t.Execute(w, r) // I put r here, because it needs 2 arguments, but I have no idea why this is ok
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:3000", nil)
}