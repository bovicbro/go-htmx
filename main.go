package main

import (
	"fmt"
	"github.com/a-h/templ"
	"htmx/components/mainContainer"
	"net/http"
)

func main() {
	http.Handle("/", templ.Handler(mainContainer.MainContainer()))

	fmt.Println("Listening on 3000")
	http.ListenAndServe(":3000", nil)
}
