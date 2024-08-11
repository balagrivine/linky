package logic

import (
	"encoding/json"
	"fmt"
	"http"
	"time"
	"database/sql"

	"github.com/go-playground/validator/v10"

	"github.com/balagrivine/linky/utils"
	"github.com/balagrivine/linky/internal/database"
)

type createUser struct {
	Email string `json:"email" validate:"required"`
}

func handleCreateUser(dbCfg *config.DBConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var param createUser

		// Decode thr request body into the param variable
		err := decoder.Decode(&param)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsong JSON: %s", err), http.StatusBadRequest)
			return
		}

		validate := validator.New()

		err := validate.Struct(param)
		if err != nil {
			http.Error(w, fmt.Sprintf("Validation failed: %s", err), http.StatusBadRequest)
			return
		}

		user, err := dbCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			Email: param.Email
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating user: %s", err), http.InternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

type createURL struct {
	OriginalURL  string `json:"original_url" validate:"required"`
	CreatedAt    string `json:"created_at" validate:"required"`
}

func CreateShortURL(apiCfg *config.DBConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		decoder := json.NewDecoder(r.Body)
		var param CreateURL

		// Decode the request body into the param variable
		err := decoder.Decode(&param)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing JSON: %s", err), http.StatusBadRequest)
			return
		}

		// Create a new validator instance
		validate := validator.New()

		err := validate.Struct(param)
		if err != nil {
			http.Error(w, fmt.Sprintf("Validation failed: %s", err), http.StatusBadRequest)
			return
		}

		currentTime := time.Now().UTC()
		nullTime := sql.NullTime{
			Time: currentTime,
			Valid: true,
		}

		url, err := apiCfg.DB.CreateURL(r.Context(), database.CreateURLParams {
			OriginalUrl:  param.OriginalURL,
			ShortUrl:     "",
			CreatedAt:    nullTime,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating URL: %s"), http.StatusInternalServerError)
			return
		}

		url_id := url.id
		// Hash the url id
		hashed_id := utils.HashURLId(url_id)
		// Encode the url id
		encoded_url, err := utils.EncodeURL(hashed_id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error encoding ID: %s", err) http.StatusInternalServerError)
			return
		}

		url.short_url := fmt.Sprintf("http://localhost/%s", encoded_url)

		json.NewEncoder(w).Encode(url)
	}
}
