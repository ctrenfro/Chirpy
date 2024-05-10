package main

import (
	"net/http"
	"sort"
	"strconv"
)

func (cfg *apiConfig) handlerUsersRetrieve(w http.ResponseWriter, r *http.Request) {
	dbUsers, err := cfg.DB.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve users")
		return
	}

	users := []User{}
	for _, dbUser := range dbUsers {
		users = append(users, User{
			ID:    dbUser.ID,
			Email: dbUser.Email,
		})
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})

	respondWithJSON(w, http.StatusOK, users)
}

func (cfg *apiConfig) handlerUsersRetrieveById(w http.ResponseWriter, r *http.Request) {
	dbUsers, err := cfg.DB.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve users")
		return
	}

	stringID := r.PathValue("userid")

	intID, err := strconv.Atoi(stringID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid ID")
		return
	}

	if intID > len(dbUsers) {
		respondWithError(w, http.StatusNotFound, "ID does not exist")
		return
	}

	user := User{
		ID:    dbUsers[intID-1].ID,
		Email: dbUsers[intID-1].Email,
	}

	respondWithJSON(w, http.StatusOK, user)
}
