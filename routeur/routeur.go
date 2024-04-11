package routeur

import (
	"log"
	"net/http"
)

func Initserv() {
	css := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))
	http.HandleFunc("/login", backend.Login)
	addr := ":8000"
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)

	}
}
