package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/schlinke/what4lunch/urlparser"
)

const (
	lunchDir = "/lunch/"
	menuDir  = "/menu/"
	port     = "8090"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>What4lunch</h1>")

	fmt.Fprintf(w, "<h1>Lunch</h1>")

	lunch := urlparser.GetLinksLunch()
	for k, v := range lunch {
		fmt.Fprintf(w, "<a href=%s>%s</a><br/>", v, k)
	}

	fmt.Fprintf(w, "<h1>Menus</h1>")
	menu := urlparser.GetLinksMenu()
	for k, v := range menu {
		fmt.Fprintf(w, "<a href=%s>%s</a><br/>", v, k)
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	log.Output(1, "Starting Webserver...")
	//urlparser.GetAll()
	urlparser.GetLinksLunch()
	r.HandleFunc("/", handler)
	r.PathPrefix(lunchDir).Handler(http.StripPrefix(lunchDir, http.FileServer(http.Dir("."+lunchDir))))
	r.PathPrefix(menuDir).Handler(http.StripPrefix(menuDir, http.FileServer(http.Dir("."+menuDir))))
	http.ListenAndServe(":"+port, r)
	log.Output(1, "Webserver started")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
