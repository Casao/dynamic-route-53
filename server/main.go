package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		addr := strings.Split(r.RemoteAddr, ":")
		w.Write([]byte(addr[0]))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
