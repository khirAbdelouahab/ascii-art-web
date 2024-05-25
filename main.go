package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	//"os"
	//"fmt"
	"ascii-art-web/banner"
)

type Page struct {
	PageTitle string
}
type Banner struct {
	TextResult string
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := Page{
		"Home Page",
	}
	t, err := template.ParseFiles("./Pages/index.html")
	if err != nil {
		log.Fatal("Error while try parsing data ", err)
	}
	t.Execute(w, data)
}

func LoadBanner(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./Pages/banner.html")
	if err != nil {
		log.Fatal("Error while try parsing data ", err)
	}

	switch r.Method {
	case "GET":
		if err != nil {
			log.Fatal("Error while try parsing data ", err)
		}
		fmt.Fprintf(w, "<script>alert('GET Method')</script>")
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Form Parse err  :0", 1)
			return
		}
		text := r.PostFormValue("text")
		newLineCounter := strings.Count(text, "\r\n")
		words := strings.Split(text, "\r\n")
		fileName :="./files/" + r.PostFormValue("banner") +".txt"
		textResult := banner.Result(words, newLineCounter, banner.ReadBannerFiles(fileName))
		result := Banner{
			TextResult: textResult,
		}
		t.Execute(w, result)
	default:
		fmt.Fprintf(w, "only get and post")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/Pages/banner.html", LoadBanner)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
