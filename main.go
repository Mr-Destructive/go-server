package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Args[1]
	fmt.Printf("Starting server at port %s\n", port)
	path := os.Args[2]
	fileServer := http.FileServer(http.Dir(path))
	http.Handle("/", fileServer)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
