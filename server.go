package main

import (
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
	fmt.Println("Serving files in the current directory on port 9001")
	http.Handle("/", &middleware{})
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
