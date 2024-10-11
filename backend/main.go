package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	const port = "8080"
	mux := http.NewServeMux()
	const filepathRoot = ".."
	//fsHandler := http.FileServer(http.Dir(filepathRoot))
	//mux.Handle("/", fsHandler)
	mux.HandleFunc("POST /api/users", apiCfg.UserCreateHandler)
	mux.HandleFunc("DELETE /api/users", apiCfg.UserDeleteHandler)
	mux.HandleFunc("PUT /api/users", apiCfg.UserUpdateHandler)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())

}
