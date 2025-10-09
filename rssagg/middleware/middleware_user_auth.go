package middleware

import (
	"fmt"
	"net/http"

	"github.com/flint92/rssagg/auth"
	"github.com/flint92/rssagg/internal/database"
	"github.com/flint92/rssagg/respod"
)

type UserAuthHandler func(w http.ResponseWriter, r *http.Request, user *database.User)

func UserAuth(queries *database.Queries, h UserAuthHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respod.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Invalid API key: %v", err))
			return
		}

		usr, err := queries.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respod.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Error getting user by api key: %v", apiKey))
			return
		}

		h(w, r, &usr)
	}
}
