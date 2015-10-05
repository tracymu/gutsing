package main

import (
	"fmt"
	"github.com/gedex/go-instagram/instagram"
	"html/template"
	"net/http"
	// "reflect"
	"strings"
	// "io/ioutil"
	// "encoding/json"
)

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("results.html")
	t.Execute(w, r) // I put r here, because it needs 2 arguments, but I have no idea why this is ok
}

func homeQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("main.html")
		t.Execute(w, r) // I put r here, because it needs 2 arguments, but I have no idea why this is ok
	} else {
		r.ParseForm()
		manipulateStuff(r.Form["food"])
	}
}

func manipulateStuff(food []string) {
	s := food
	s = strings.Split(s[0], " ")
	callInstagram(strings.Join(s, ""))
}

func callInstagram(query string) {
	client := instagram.NewClient(nil)

	media, _, _ := client.Tags.RecentMedia(query, nil)
	pics := [20]string{}
	index := 0
	for index < len(media) {
		pics[index] = media[index].Images.LowResolution.URL
		index += 1
	}
}

func main() {
	http.HandleFunc("/results", resultsHandler)
	http.HandleFunc("/", homeQuery)
	http.ListenAndServe("localhost:3000", nil)
}