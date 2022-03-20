package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	serveHttp()
}

func serveHttp() {
	serviceHost, ok := os.LookupEnv("SERVICE_HOST")

	if !ok {
		serviceHost = ":80"
	}

	fs := http.FileServer(http.Dir("/opt/mwcertbot/.well-known/"))
	http.Handle("/.well-known/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hallo Welt"))
	})

	fmt.Printf("Listening for incomming http-requests on %s\n", serviceHost)

	if err := http.ListenAndServe(serviceHost, nil); err != nil {
		log.Fatal(err)
	}

}
