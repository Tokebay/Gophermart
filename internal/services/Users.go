package service

import (
	"encoding/json"
	"net/http"

	"github.com/Tokebay/yandex-diplom/internal/api"
	"github.com/Tokebay/yandex-diplom/internal/logger"

	"github.com/Tokebay/yandex-diplom/internal/models"
)

type UserHandler struct {
	users *models.User
}

type UserInput struct {
	Login    string `json:"login" validate:"required,gte=2"`
	Password string `json:"password" validate:"required,gte=4"`
}

func NewUserHandler(users *models.User) *UserHandler {
	return &UserHandler{
		users: users,
	}
}

func (s *api.Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var userInput UserInput
	if err := parseJSONBody(r, &userInput); err != nil {
		response.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if err := h.users.SignUp(userInput); err != nil {
		response.Error(w, "Failed to register user", http.StatusInternalServerError)
		logger.Log.Error("Error registering user", err)
		return
	}

	response.JSON(w, "User registered successfully", http.StatusOK)
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userInput UserInput
	if err := parseJSONBody(r, &userInput); err != nil {
		response.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// token, err := h.users.SignIn(userInput)
	// if err != nil {
	// 	if err == domain.ErrUserNotFound {
	// 		response.Error(w, "Invalid username or password", http.StatusUnauthorized)
	// 	} else {
	// 		response.Error(w, "Failed to authenticate user", http.StatusInternalServerError)
	// 		logger.Log.Error("Error authenticating user", err)
	// 	}
	// 	return
	// }

	// response.JSON(w, map[string]string{"token": token}, http.StatusOK)
}

func parseJSONBody(r *http.Request, target interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}
