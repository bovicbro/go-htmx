package main

import (
	"fmt"
	"github.com/a-h/templ"
	"htmx/components/mainContainer"
	"net/http"
)

func main() {
	mc := mainContainer.New()
	thing := templ.Handler(mc.CreateCmp())
	http.Handle("/", thing)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Listening on 3000")
	http.ListenAndServe(":3000", nil)
}
