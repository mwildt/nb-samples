package main

import (
	"log"
	"os"
	"net/http"
)

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {

	log.Printf("Running NB Sample Web-Application")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", requestLogger(fs))

	serviceHost, ok := os.LookupEnv("SERVICE_HOST")

	if !ok {
		serviceHost = ":80"
	}

	log.Printf("Listening on %s\n", serviceHost)
	
	if err := http.ListenAndServe(serviceHost, nil); err != nil {
		log.Fatal(err)
	}

}