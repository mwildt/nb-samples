package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {

	log.Printf("Running NB Sample Web-Application")

	serviceMode, ok := os.LookupEnv("SERVICE_MODE")

	if !ok {
		serviceMode = "HTTP"
	}

	switch strings.ToLower(serviceMode) {

	case "http":
		serveHttp()
	case "tcp":
		serveTcp()
	}

}

func serveTcp() {
	serviceHost, ok := os.LookupEnv("SERVICE_HOST")

	if !ok {
		serviceHost = ":3300"
	}

	listener, err := net.Listen("tcp", serviceHost)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening for incomming tcp-Connections on %s\n", serviceHost)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Accept Error", err)
			continue
		}

		fmt.Printf("Accepted Connection from %s\n", conn.RemoteAddr())

		go func(conn net.Conn) {

			defer conn.Close()

			buffer := make([]byte, 1024)
			receivedBytes, err := conn.Read(buffer)
			message := buffer[:receivedBytes]
			fmt.Printf("Incomming Request Payload '%s' (%v bytes)", message, receivedBytes)

			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("Sending Response to Client\n")
				conn.Write([]byte("Hello "))
				conn.Write([]byte(message))
			}
		}(conn)
	}

}

func serveHttp() {
	serviceHost, ok := os.LookupEnv("SERVICE_HOST")

	if !ok {
		serviceHost = ":80"
	}

	echoHost, ok := os.LookupEnv("ECHO_HOST")

	if !ok {
		echoHost = "127.0.0.1:3300"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s", r.Method, r.URL.Path)

		if conn, err := net.Dial("tcp", echoHost); err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Fehler beim Aufbau der TCP-Connection"))
		} else {

			fmt.Printf("Send tcp client Request \n")
			fmt.Fprintf(conn, "Max Mustermann")

			buffer := make([]byte, 1024)
			if receivedBytes, err := conn.Read(buffer); err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte("Fehler beim Abrufen der Tcp-Daten"))
			} else {
				message := buffer[:receivedBytes]
				fmt.Printf("Message '%s' (recieved %d bytes)\n", message, receivedBytes)
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte(message))
			}
		}
	})

	fmt.Printf("Listening for incomming http-requests on %s\n", serviceHost)

	if err := http.ListenAndServe(serviceHost, nil); err != nil {
		log.Fatal(err)
	}

}
