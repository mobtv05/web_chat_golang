package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var (
	tmpl = template.Must(template.ParseFiles("templates/index2.html"))
	text string
)

func hendler(w http.ResponseWriter, r *http.Request) {
	log.Println("Client conected: " + r.RemoteAddr)
	defer log.Println("Client disconected: " + r.RemoteAddr)
	text = text + "<br>" + r.FormValue("message")
	tmpl.Execute(w, text)
	io.WriteString(w, text)

}

func main() {
	http.HandleFunc("/", hendler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
