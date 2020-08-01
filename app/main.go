package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/clozed2u/test-lightnetio/calculator"
)

type payload struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type errResp struct {
	Error string `json:"error"`
}

type resultResp struct {
	Result float64 `json:"result"`
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		p := payload{}
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			resp := errResp{Error: "bad request"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp)
		} else {
			resp := resultResp{Result: calculator.Sum(p.A, p.B)}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
		}
	} else {
		resp := errResp{Error: "not implement"}
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(resp)
	}
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		p := payload{}
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			resp := errResp{Error: "bad request"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp)
		} else {
			resp := resultResp{Result: calculator.Sub(p.A, p.B)}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
		}
	} else {
		resp := errResp{Error: "not implement"}
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(resp)
	}
}

func mulHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		p := payload{}
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			resp := errResp{Error: "bad request"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp)
		} else {
			resp := resultResp{Result: calculator.Mul(p.A, p.B)}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
		}
	} else {
		resp := errResp{Error: "not implement"}
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(resp)
	}
}

func divHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		p := payload{}
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			resp := errResp{Error: "bad request"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp)
		} else {
			result, err := calculator.Div(p.A, p.B)
			if err != nil {
				resp := errResp{Error: err.Error()}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(resp)
			} else {
				resp := resultResp{Result: result}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(resp)
			}
		}
	} else {
		resp := errResp{Error: "not implement"}
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(resp)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	resp := errResp{Error: "not found"}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	PORT := os.Getenv("APP_PORT")
	if PORT == "" {
		log.Fatalln("APP_PORT environment variable is not set")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/calculator.sum", sumHandler)
	mux.HandleFunc("/calculator.sub", subHandler)
	mux.HandleFunc("/calculator.mul", mulHandler)
	mux.HandleFunc("/calculator.div", divHandler)
	mux.HandleFunc("/", notFoundHandler)
	srv := &http.Server{Addr: fmt.Sprintf(":%s", PORT), Handler: mux}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err.Error())
		}
	}()
	log.Println(fmt.Sprintf("App server running at :%s", PORT))
	<-done
	log.Println("App server stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("App server shutdown failed: %s\n", err.Error())
	}
	log.Println("App server exited properly")
}
