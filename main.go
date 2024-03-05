package main

import (
	"fmt"
	"github.com/a-h/templ"
	"htmx/components/mainContainer"
	"net/http"
)

func main() {
	mc := mainContainer.New()
	http.Handle("/", templ.Handler(mc))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Listening on 3000")
	http.ListenAndServe(":3000", nil)
}
