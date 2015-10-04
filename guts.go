package main

import (
	"fmt"
	"github.com/gedex/go-instagram/instagram"
	"html/template"
	"net/http"
	"reflect"
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

	// s := []string{"foo", "bar", "baz"}
	// fmt.Println(s)
	s = strings.Split(s[0], " ")
	// fmt.Println(strings.Join(s, ""))
	callInstagram(strings.Join(s, ""))
}

func callInstagram(query string) {
	client := instagram.NewClient(nil)

	media, next, err := client.Tags.RecentMedia(query, nil)
	// media, next, err := client.Users.RecentMedia("3", nil)
	fmt.Println(reflect.TypeOf(media))
	fmt.Println(next)
	fmt.Println(err)
	// I guess if don't want to use next and err, use _ in the assignment line.
}

type Tag struct {
	MediaCount int    `json:"media_count,omitempty"`
	Name       string `json:"name,omitempty"`
}

// func (c *Client) NewRequest('method', urlStr string, body string) (*http.Request, error)

// func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error)

func main() {
	http.HandleFunc("/results", resultsHandler)
	http.HandleFunc("/", homeQuery)
	http.ListenAndServe("localhost:3000", nil)
}

// what are locations vs geographies?
//  https://godoc.org/github.com/gedex/go-instagram/instagram
