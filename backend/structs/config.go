package structs

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
func (cfg *ApiConfig) UserLogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	})
	jsonResponse.RespondWithJSON(w, 200, "ok")
}
func (cfg *ApiConfig) UserProfileHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Header)

	cookie, err := r.Cookie("token")
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusUnauthorized, "Couldn't find JWT in cookies", err)
		return
	}

	token := cookie.Value
	fmt.Println("JWT Token from cookie:", token)

	userID, err := auth.ValidateJWT(token, cfg.SecretKey)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}
	user, err := cfg.DB.GetUserById(context.Background(), userID)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "email or password wrong", err)
		return
	}
	jsonResponse.RespondWithJSON(w, 200, database.User{
		ID:    user.ID,
		Email: user.Email,
	})

}

func (cfg *ApiConfig) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type response struct {
		database.User
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	user, err := cfg.DB.GetUser(context.Background(), params.Email)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "email or password wrong", err)
		return
	}
	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "email or password wrong", err)
		return
	}
	accessToken, err := auth.MakeJWT(
		user.ID,
		cfg.SecretKey,
		time.Hour,
	)
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Couldn't create access JWT", err)
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Couldn't create refresh token", err)
		return
	}

	_, err = cfg.DB.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: sql.NullTime{Time: time.Now().AddDate(0, 0, 60), Valid: true},
	})
	if err != nil {
		jsonResponse.RespondWithError(w, http.StatusInternalServerError, "Couldn't save refresh token", err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    accessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	})
	jsonResponse.RespondWithJSON(w, 200, response{
		User: database.User{
			ID:    user.ID,
			Email: user.Email,
		},
		Token:        accessToken,
		RefreshToken: refreshToken,
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
