package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	tmpl = template.Must(template.ParseFiles("templates/index2.html"))
	text string
)

func hendler(w http.ResponseWriter, r *http.Request) {
	//log.Println("Client conected: " + r.RemoteAddr)
	ip := strings.Split(r.RemoteAddr, ":")[0]
	//defer log.Println("Client disconected: " + r.RemoteAddr)
	text = text + "<br>" + r.FormValue("message")
	tmpl.Execute(w, text)
	io.WriteString(w, text)
	fmt.Println("Message from: ", ip, " and the text of this message", text)
	// w.Write([]byte(response))

}

func main() {
	http.HandleFunc("/", hendler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
