package main

import (
	"net/http"

	"github.com/flint92/rssagg/respod"
)

func handlerErr(w http.ResponseWriter, _ *http.Request) {
	respod.RespondWithError(w, http.StatusBadRequest, "Something went wrong")
}
