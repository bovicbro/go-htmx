package main

import (
	"context"
	"fmt"
	"htmx/components/mainContainer"
	"net/http"
)

func main() {
	mc := mainContainer.New()
	http.HandleFunc("/", func(w http.ResponseWriter, res *http.Request) {
		mc.CreateCmp().Render(context.Background(), w)
	})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Listening on 3000")
	http.ListenAndServe(":3000", nil)
}
