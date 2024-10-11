package structs

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Can't create user", err)
		return
	}
	user, err := cfg.DB.CreateUser(context.Background(), database.CreateUserParams{
		HashedPassword: hashed,
		Email:          params.Email,
	})
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Can't create user", err)
		return
	}

	jsonResponse.RespondWithJSON(w, 201, database.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}

func (cfg *ApiConfig) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID uuid.UUID `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	err = cfg.DB.DeleteUserById(context.Background(), params.ID)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Cant find user", err)
		return
	}

	jsonResponse.RespondWithJSON(w, http.StatusNoContent, "Successfully deleted user")

}
func (cfg *ApiConfig) UserUpdateEmailHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string    `json:"email"`
		ID    uuid.UUID `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	user, err := cfg.DB.UpdateUserEmailById(context.Background(), database.UpdateUserEmailByIdParams{
		Email: params.Email,
		ID:    params.ID,
	})
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Cant find user", err)
		return
	}

	jsonResponse.RespondWithJSON(w, http.StatusAccepted, database.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})

}

func (cfg *ApiConfig) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query().Get("action")
	type parameters struct {
		Password string    `json:"password"`
		Email    string    `json:"email"`
		ID       uuid.UUID `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	if action == "password" {
		hashed, err := auth.HashPassword(params.Password)
		if err != nil {
			jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Can't create user", err)
			return
		}
		_, err = cfg.DB.UpdateUserPasswordById(context.Background(), database.UpdateUserPasswordByIdParams{
			HashedPassword: hashed,
			ID:             params.ID,
		})
		if err != nil {
			jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Cant find user", err)
			return
		}

		jsonResponse.RespondWithJSON(w, http.StatusAccepted, "passwordUpdated")
		return
	} else if action == "email" {
		user, err := cfg.DB.UpdateUserEmailById(context.Background(), database.UpdateUserEmailByIdParams{
			Email: params.Email,
			ID:    params.ID,
		})
		if err != nil {
			jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Cant find user", err)
			return
		}

		jsonResponse.RespondWithJSON(w, http.StatusAccepted, database.User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		})
		return
	}
	jsonResponse.RespondWithError(w, http.StatusBadRequest, "No such action", err)

}
