package structs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xristoskrik/yourmusic/internal/database"
)

type ApiConfig struct {
	DB        *database.Queries
	SecretKey string
}

func (cfg *ApiConfig) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	user, err := cfg.DB.CreateUser(context.Background(), database.CreateUserParams{
		HashedPassword: params.Password,
		Email:          params.Email,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	respondWithJSON(w, 201, database.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
