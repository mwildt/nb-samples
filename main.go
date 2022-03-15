package main

import (
	"log"
	"os"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	serviceHost, ok := os.LookupEnv("SERVICE_HOST")

	if !ok {
		serviceHost = ":80"
	}

	log.Printf("Listening on %s\n", serviceHost)
	
	if err := http.ListenAndServe(serviceHost, nil); err != nil {
		log.Fatal(err)
	}

}