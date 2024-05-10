package main

import (
	"net/http"
	"sort"
	"strconv"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:   dbChirp.ID,
			Body: dbChirp.Body,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].ID < chirps[j].ID
	})

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpsRetrieveById(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
		return
	}

	stringID := r.PathValue("chirpid")

	intID, err := strconv.Atoi(stringID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid ID")
		return
	}

	if intID > len(dbChirps) {
		respondWithError(w, http.StatusNotFound, "ID does not exist")
		return
	}

	chirp := Chirp{
		ID:   dbChirps[intID-1].ID,
		Body: dbChirps[intID-1].Body,
	}

	respondWithJSON(w, http.StatusOK, chirp)
}
