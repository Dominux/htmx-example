package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("running server on http://localhost:8000")

	http.HandleFunc("/", filmsPage)
	http.HandleFunc("/add-film/", addFilm)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func filmsPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{
				Title:    "The Shawshank Redemption",
				Director: "Frank Darabont",
			},
			{
				Title:    "The Godfather",
				Director: "Francis Ford Coppola",
			},
			{
				Title:    "The Dark Knight",
				Director: "Christopher Nolan",
			},
			{
				Title:    "12 Angry Men",
				Director: "Sidney Lumet",
			},
			{
				Title:    "Schindler's List",
				Director: "Steven Spielberg",
			},
		},
	}
	tmpl.Execute(w, films)
}

func addFilm(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)

	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}
