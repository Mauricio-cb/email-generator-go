package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Mauricio-cb/email-generator-go/Encrypt"
)

var tmpl *template.Template

type Data struct {
	Email    string
	Password string
}

func emailGeneratorHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Email:    Encrypt.RandomEmail(),
		Password: Encrypt.EncryptPassword("daz"),
	}

	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()

	tmpl = template.Must(template.ParseFiles("./static/html/email.html"))
	fs := http.FileServer(http.Dir("./static/"))
	
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", emailGeneratorHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))



}
