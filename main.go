package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var port string
	var path string
	params := os.Args
	if len(params) > 1 {
		port = os.Args[1]
	} else {
		port = "8001"
	}
	if len(params) > 2 {
		path = os.Args[2]
	} else {
		path = "."
	}
	fmt.Printf("Starting server at port %s\n", port)
	fileServer := http.FileServer(http.Dir(path))
	http.Handle("/", fileServer)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
