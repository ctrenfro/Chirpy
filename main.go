package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)
	fs := http.FileServer(http.Dir("."))
	metrics := &apiConfig{}

	mux.Handle("/app/*", metrics.middlewareMetricsInc(http.StripPrefix("/app", fs)))
	mux.HandleFunc("GET /api/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

	})
	mux.Handle("GET /api/metrics", metrics.serve())
	mux.Handle("GET /api/reset", metrics.reset())
	mux.Handle("GET /admin/metrics", metrics.admin())
	mux.Handle("POST /api/validate_chirp", validate)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: corsMux,
	}

	server.ListenAndServe()

}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
