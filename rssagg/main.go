package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/flint92/rssagg/feed"
	"github.com/flint92/rssagg/internal/database"
	"github.com/flint92/rssagg/middleware"
	"github.com/flint92/rssagg/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("$DATABASE_URL must be set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	queries := database.New(conn)

	userClient := user.NewClient(queries)
	feedClient := feed.NewClient(queries)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerHealth)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", userClient.CreateUser)
	v1Router.Get("/users", middleware.UserAuth(queries, userClient.GetUser))
	v1Router.Post("/feeds", middleware.UserAuth(queries, feedClient.CreateFeed))

	router.Mount("/api/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server starting on port %s", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
