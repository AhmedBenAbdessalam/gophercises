package main

import (
	"fmt"
	"net/http"

	urlshort "github.com/AhmedBenAbdessalam/gophercises/urlshortener"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	_ = mapHandler
	for i, v := range pathsToUrls {
		mux.Handle(i, http.RedirectHandler(v, http.StatusSeeOther))
	}
	/*
			// Build the YAMLHandler using the mapHandler as the
			// fallback
			yaml := `
		- path: /urlshort
		  url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		  url: https://github.com/gophercises/urlshort/tree/solution
		`
			yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
			if err != nil {
				panic(err)
			}
	*/
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
	//http.ListenAndServe(":8080", mux)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
