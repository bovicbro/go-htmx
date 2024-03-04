package main

import (
	"fmt"
	"github.com/a-h/templ"
	"htmx/components/mainContainer"
	"net/http"
)

func main() {
	mainContainer := mainContainer.MainContainer()
	http.Handle("/", templ.Handler(mainContainer))
	http.HandleFunc("/stuff", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there")
	})
	fmt.Println("Listening on 3000")
	http.ListenAndServe(":3000", nil)
}
