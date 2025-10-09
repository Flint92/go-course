package feed

import (
	"time"

	"github.com/flint92/rssagg/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"userId"`
}

func databaseFeedToFeed(feed database.Feed) *Feed {
	return &Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}
