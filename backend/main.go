package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/xristoskrik/yourmusic/internal/database"
	"github.com/xristoskrik/yourmusic/structs"
)

func main() {

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	secret := os.Getenv("SECRET_KEY")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	dbQueries := database.New(db)
	apiCfg := structs.ApiConfig{
		DB:        dbQueries,
		SecretKey: secret,
	}
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Set the specific origin you allow
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // This allows cookies and other credentials to be sent
	})
	r.Use(cors.Handler)

	r.Route("/api", func(r chi.Router) {
		r.Post("/users", apiCfg.UserCreateHandler)
		r.Delete("/users", apiCfg.UserDeleteHandler)
		r.Put("/users", apiCfg.UserUpdateHandler)
		r.Post("/users/login", apiCfg.LoginUserHandler)
		r.Get("/users/profile", apiCfg.UserProfileHandler)
		r.Post("/users/logout", apiCfg.UserLogoutHandler)

	})
	const port = "8080"
	const filepathRoot = ".." // Your file server's root directory
	fsHandler := http.FileServer(http.Dir(filepathRoot))
	r.Handle("/*", fsHandler)
	//const filepathRoot = ".."
	//fsHandler := http.FileServer(http.Dir(filepathRoot))
	//mux.Handle("/", fsHandler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())

}
