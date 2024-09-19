package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/igordonatti/REST-API/models"
	"github.com/igordonatti/REST-API/utils"
)

var validate *validator.Validate

// / Put this in one separate archive
type UserInput struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []models.User
	models.DB.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var input UserInput

	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	user := &models.User{
		Name:  input.Name,
		Email: input.Email,
	}

	models.DB.Create(user)

	w.Header().Set("Content-Type", "application/json") // teria como fazer uma forma melhor isso aqui tem nao

	json.NewEncoder(w).Encode(user)
}
