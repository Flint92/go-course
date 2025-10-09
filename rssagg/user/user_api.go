package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/flint92/rssagg/auth"
	"github.com/flint92/rssagg/internal/database"
	"github.com/flint92/rssagg/respod"
	"github.com/google/uuid"
)

type ApiConfig struct {
	DB *database.Queries
}

func NewApiConfig(db *database.Queries) *ApiConfig {
	return &ApiConfig{DB: db}
}

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	var p parameters
	err := decoder.Decode(&p)
	if err != nil {
		respod.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	defer func(c io.ReadCloser) {
		err := c.Close()
		if err != nil {
			log.Printf("Error closing body: %v", err)
		}
	}(r.Body)

	usr, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      p.Name,
	})
	if err != nil {
		respod.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respod.RespondWithJSON(w, http.StatusCreated, databaseUserToUser(usr))
}

func (apiCfg *ApiConfig) GetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respod.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Invalid API key: %v", err))
		return
	}

	usr, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respod.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Error getting user by api key: %v", apiKey))
		return
	}

	respod.RespondWithJSON(w, http.StatusOK, databaseUserToUser(usr))
}
