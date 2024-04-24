package main

import (
	"context"
	"fmt"
	"htmx/components/mainContainer"
	"net/http"
)

func main() {
	mc := mainContainer.New()
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		mc.CreateCmp().Render(context.Background(), w)
	})
	fmt.Println("Listening on 3000")
	http.ListenAndServe(":3000", nil)
}
