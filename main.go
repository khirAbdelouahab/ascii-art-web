package main

import (
	//"os"
	//"fmt"
	"ascii-art-web/banner"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Page struct {
	PageTitle string
}
type Banner struct {
	TextResult string
}

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w,"<script>alert(1)</script>")
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
			fmt.Fprintf(w, "Form Parse err  : %v", err)
			return
		}
		text := r.PostFormValue("text")
		newLineCounter := strings.Count(text,"\\n")
		words := strings.Split(text,"\\n")
		textResult := banner.Result(words,newLineCounter,banner.ReadBannerFiles("./files/standard.txt"))
		result := Banner{
			TextResult: textResult ,
		}
		t.Execute(w, result)
	default:
		fmt.Fprintf(w, "only get and post")
	}
}
func main() {
	/*sweaters := Inventory{"wool", 17}
	s := "{{.Count -}} items are made of {{- .Material}} mybee\n"
	tmpl, err := template.New("test").Parse(s)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}*/
	/*text := "hello"
	newLineCounter := strings.Count(text,"\\n")
	words := strings.Split(text,"\\n")
	textResult := banner.Result(words,newLineCounter,banner.ReadBannerFiles("./files/standard.txt"))
	fmt.Println(textResult)*/
	http.HandleFunc("/", handler)
	http.HandleFunc("/Pages/banner.html", LoadBanner)
	http.ListenAndServe(":8080", nil)
}
