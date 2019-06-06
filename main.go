package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func main() {
	var message string
	tmpl := template.Must(template.ParseFiles("template.html"))
	// msg := template.Must(template.ParseFiles("message.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(path.Dir(r.URL.Path))
		var err error
		switch path.Dir(r.URL.Path) {
		case "/":
			err = tmpl.Execute(w, message)
			fmt.Println("root")
		case "/submit":
			message = path.Base(r.URL.Path)
			err = tmpl.ExecuteTemplate(w, "batch", message)
			fmt.Println("message", message)
		}
		if err != nil {
			log.Fatalln(err)
		}

	})

	http.ListenAndServe(":8080", nil)
}
