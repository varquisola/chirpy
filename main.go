package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	handler := http.NewServeMux()
	port := "8080"
	filePathRoot := "."

	fmt.Println("Server starting on port", port)

	apiCfg := &apiConfig{
		fileserverHits: atomic.Int32{},
	}

	// Handle /app/ prefix
	handler.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filePathRoot)))))
	handler.HandleFunc("/healthz", heartBeat)
	handler.HandleFunc("/metrics", apiCfg.printMetrics)
	handler.HandleFunc("/reset", apiCfg.resetMetrics)

	server := &http.Server{
		Handler: handler,
		Addr:    ":" + port,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error with the server")
		return
	}

}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) printMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())))
}
