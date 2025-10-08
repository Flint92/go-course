package main

import "net/http"

func handlerHealth(w http.ResponseWriter, _ *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"health": "UP"})
}
