package feed

import (
	"fmt"
	"net/http"
	"time"

	"github.com/flint92/rssagg/internal/database"
	"github.com/flint92/rssagg/req"
	"github.com/flint92/rssagg/respod"
	"github.com/google/uuid"
)

type Client struct {
	DB *database.Queries
}

func NewClient(db *database.Queries) *Client {
	return &Client{DB: db}
}

func (feedClient *Client) CreateFeed(w http.ResponseWriter, r *http.Request, user *database.User) {

	type parameters struct {
		Name   string    `json:"name"`
		Url    string    `json:"url"`
		UserID uuid.UUID `json:"userID"`
	}

	var p parameters
	err := req.ReadToJson(r, &p)
	if err != nil {
		respod.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      p.Name,
		Url:       p.Url,
		UserID:    user.ID,
	}

	createdFeed, err := feedClient.DB.CreateFeed(r.Context(), feed)
	if err != nil {
		respod.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	respod.RespondWithJSON(w, http.StatusCreated, databaseFeedToFeed(createdFeed))
}
