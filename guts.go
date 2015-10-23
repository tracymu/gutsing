package main

import (
	"fmt"
	"github.com/gedex/go-instagram/instagram"
	"html/template"
	"net/http"
	"strings"
	"math/rand"
)

var pics = [20]string{}

type Pic struct {
	UrlArray string
}

func homeQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("main.html")
		t.Execute(w, r) 
	} else {
		r.ParseForm()
		manipulateStuff(r.Form["food"])
	}
	http.Redirect(w,r,"/results", http.StatusFound)
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {	
	random_pic := pics[rand.Intn(len(pics))]
	result_pics := Pic{random_pic}
	template, _ := template.ParseFiles("results.html")
	template.Execute(w, result_pics)

	fmt.Println("here in results handlre fucntion")
	fmt.Println(result_pics.UrlArray[0])

}


func manipulateStuff(food []string) {
	s := food
	s = strings.Split(s[0], " ")
	callInstagram(strings.Join(s, ""))
}

func callInstagram(query string) {
	client := instagram.NewClient(nil)
	media, _, _ := client.Tags.RecentMedia(query, nil)

	index := 0
	for index < len(media) {
		pics[index] = media[index].Images.LowResolution.URL
		index += 1
	}
  return
}

func main() {
	http.HandleFunc("/results", resultsHandler)
	http.HandleFunc("/", homeQuery)
	http.ListenAndServe("localhost:3000", nil)

}