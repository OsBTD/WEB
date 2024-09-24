package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"webb/ascii"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintln(w, "error hh 404")
		return
	}
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, err)
}

func Ascii(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("text")
	banner := r.FormValue("banner")
	Print := ascii.PrintArt(input, banner)
	fmt.Println(w, input, banner)
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, Print)
}

func main() {
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	imageFs := http.FileServer(http.Dir(filepath.Join("templates", "images")))
	http.Handle("/images/", http.StripPrefix("/images/", imageFs))

	http.HandleFunc("/", Home)
	http.HandleFunc("/ascii", Ascii)
	log.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
