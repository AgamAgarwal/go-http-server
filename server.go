package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type middleware struct{}

func (*middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL, "from", r.RemoteAddr)
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
}

func main() {
	portPtr := flag.Int("port", 8000, "Port to bind")

	// Parse all flags
	flag.Parse()

	fmt.Printf("Serving files in the current directory on port %d\n", *portPtr)
	http.Handle("/", &middleware{})
	err := http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
