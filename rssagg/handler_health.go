package main

import (
	"net/http"

	"github.com/flint92/rssagg/respod"
)

func handlerHealth(w http.ResponseWriter, _ *http.Request) {
	respod.RespondWithJSON(w, http.StatusOK, map[string]string{"health": "UP"})
}
