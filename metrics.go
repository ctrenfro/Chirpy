package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits int32
}

const adminPage = `<html>
<body>
  <h1>Welcome, Chirpy Admin</h1>
  <p>Chirpy has been visited %d times!</p>
</body>
</html>`

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32((&cfg.fileserverHits), 1)

		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) serve() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "Hits: %d\n", atomic.LoadInt32(&cfg.fileserverHits))
	})
}
func (cfg *apiConfig) admin() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, adminPage, atomic.LoadInt32(&cfg.fileserverHits))
	})
}

func (cfg *apiConfig) reset() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		atomic.SwapInt32(&cfg.fileserverHits, 0)
		_, _ = fmt.Fprintln(w, "Hits: 0")
	})
}
