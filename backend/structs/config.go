package structs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xristoskrik/yourmusic/auth"
	"github.com/xristoskrik/yourmusic/internal/database"
	jsonResponse "github.com/xristoskrik/yourmusic/json"
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
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	hashed, err := auth.HashPassword(params.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err := cfg.DB.CreateUser(context.Background(), database.CreateUserParams{
		HashedPassword: hashed,
		Email:          params.Email,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonResponse.RespondWithJSON(w, 201, database.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
