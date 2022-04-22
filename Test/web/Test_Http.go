package web

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	s := "<html><body><H1>Hello, world!</H1></body></html>"
	_, _ = fmt.Fprintf(w, "%s", s)
	log.Printf("%s", s)
}

func Start() {
	fmt.Println("server start.")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
