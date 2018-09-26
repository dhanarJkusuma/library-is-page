package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("/apps/projects/goprojects/src/library-is/templates/*.gohtml"))
}

func main() {

	fs := http.FileServer(http.Dir("/apps/projects/goprojects/src/library-is/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handleIndex)
	fmt.Println("[Library IS Page -  App] Running ...")
	http.ListenAndServe(":3013", nil)
}

func handleIndex(wr http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(wr, "index.gohtml", nil)
}
