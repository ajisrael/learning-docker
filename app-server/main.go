package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	m := http.NewServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	addr := ":" + port

	m.HandleFunc("/", handlePage)

	svr := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on port ", addr)
	err := svr.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	const page = `
		<html>
		<body>
		<p> Hello from Docker! I'm a Go server. </p>
		</body>
		</html>
	`

	w.WriteHeader(200)
	w.Write([]byte(page))
}
