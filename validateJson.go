package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func validate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type errorMsg struct {
		Error string `json:"error"`
	}

	type successMsg struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respBody := errorMsg{
			Error: "Something went wrong",
		}
		dat, err := json.Marshal(respBody)
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		w.Write(dat)
		return
	}

	if len(params.Body) > 140 {
		respBody := errorMsg{
			Error: "Chirp is too long",
		}
		dat, err := json.Marshal(respBody)
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(400)
		w.Write(dat)
		return
	}

	respBody := successMsg{
		Valid: true,
	}
	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
