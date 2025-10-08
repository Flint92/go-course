package main

import "net/http"

func handlerErr(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, http.StatusBadRequest, "Something went wrong")
}
