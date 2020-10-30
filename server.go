package cycapi

import (
	"log"
	"net/http"
)

// StartServer starts a simple http server
func StartServer(address string, handler http.Handler) {
	log.Println("Serving Connections On", address);
	if err := http.ListenAndServe(address, handler); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

// StartTLSServer starts a TLS server with provided TLS cert and key files
func StartTLSServer(address string, handler http.Handler, certFile, keyFile string) {
	if err := http.ListenAndServeTLS(address, certFile, keyFile, handler); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTPS server ListenAndServe: %v", err)
	}
}