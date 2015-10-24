package main

import (
	"fmt"
	"github.com/gedex/go-instagram/instagram"
	"html/template"
	"net/http"
	"strings"
)

var pics = [20]string{}

type Pic struct {
	Images template.HTML
}

func homeQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("main.html")
		t.Execute(w, r) 
	} else {
		r.ParseForm()
		manipulateStuff(r.Form["food"])
		http.Redirect(w,r,"/results", http.StatusFound)
	}
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {	
	html := ""

  for _, pic := range pics {
 	  html += fmt.Sprintf("<div><img src='%v' /></div>", pic)
   }

  result_pics := Pic{template.HTML(html)}
	template, _ := template.ParseFiles("results.html")
	template.Execute(w, result_pics)
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