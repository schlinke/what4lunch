package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/schlinke/what4lunch/dbaccess"
	"github.com/schlinke/what4lunch/urlparser"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>What4lunch</h1>")
	lunch := dbaccess.GetLunch()

	fmt.Fprintf(w, "<h1>Lunch</h1>")
	for k, v := range lunch {
		fmt.Fprintf(w, "<a href=%s>%s</a><br/>", urlparser.ParseURL(v, time.Now()), k)
	}

	fmt.Fprintf(w, "<h1>Menus</h1>")
	menu := dbaccess.GetMenu()
	for k, v := range menu {
		fmt.Fprintf(w, "<a href=%s>%s</a><br/>", urlparser.ParseURL(v, time.Now()), k)
	}
}

func main() {
	log.Output(1, "Starting Webserver...")
	http.HandleFunc("/", handler)
	log.Output(1, "Webserver started")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
