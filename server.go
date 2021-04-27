package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var timeFormat = "Jan 2 2006 15:04:05"

type middleware struct{
	dir string
}

func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s from %s at %s\n", r.Method, r.URL, r.RemoteAddr,
		time.Now().Format(timeFormat))
	http.FileServer(http.Dir(m.dir)).ServeHTTP(w, r)
}

func main() {
	portPtr := flag.Int("port", 8000, "Port to bind")
	dir := flag.String("dir", ".", "Directory to serve")

	// Parse all flags
	flag.Parse()

	fmt.Printf("Serving files in the directory %s on port %d\n", *dir, *portPtr)
	http.Handle("/", &middleware{dir: *dir})
	err := http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
