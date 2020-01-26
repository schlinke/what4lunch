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
	menu := dbaccess.ConnectAndGet()

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
