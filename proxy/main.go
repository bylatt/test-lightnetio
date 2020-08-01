package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type errResp struct {
	Error string `json:"error"`
}

func isPathAllow(path string) bool {
	allowPaths := map[string]bool{"/calculator.sum": true, "/calculator.sub": true, "/calculator.mul": true, "/calculator.div": true}
	if val, ok := allowPaths[path]; ok {
		return val
	}
	return false
}

func main() {
	PORT := os.Getenv("PROXY_PORT")
	UPSTREAM := os.Getenv("PROXY_UPSTREAM")
	if PORT == "" {
		log.Fatalln("PROXY_PORT environment variable is not set")
	}
	if UPSTREAM == "" {
		log.Fatalln("PROXY_UPSTREAM environment variable is not set")
	}
	proxyURL, err := url.Parse(UPSTREAM)
	if err != nil {
		log.Fatalln("Failed to parse proxy upstream")
	}
	proxy := httputil.NewSingleHostReverseProxy(proxyURL)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if isPathAllow(r.URL.Path) {
			proxy.ServeHTTP(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			resp := errResp{Error: "endpoint is not allowed"}
			json.NewEncoder(w).Encode(resp)
		}
	})
	srv := &http.Server{Addr: fmt.Sprintf(":%s", PORT), Handler: mux}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err.Error())
		}
	}()
	log.Println(fmt.Sprintf("Proxy server running at :%s", PORT))
	<-done
	log.Println("Proxy server stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Proxy server shutdown failed: %s\n", err.Error())
	}
	log.Println("Proxy server exited properly")
}
