package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func clientIP(r *http.Request) string {
	// If behind proxy (Coolify/Traefik), this is typically set.
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// Can be "client, proxy1, proxy2"
		parts := strings.Split(xff, ",")
		if len(parts) > 0 {
			ip := strings.TrimSpace(parts[0])
			if ip != "" {
				return ip
			}
		}
	}
	if xrip := strings.TrimSpace(r.Header.Get("X-Real-Ip")); xrip != "" {
		return xrip
	}

	// Fallback
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && host != "" {
		return host
	}
	return r.RemoteAddr
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := clientIP(r)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, ip)
	})

	addr := ":8080"
	log.Fatal(http.ListenAndServe(addr, mux))
}
