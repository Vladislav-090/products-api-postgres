package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"product-api-postgres/internal/auth"
	"product-api-postgres/internal/models"
	"product-api-postgres/internal/response"
	"product-api-postgres/internal/storage"
)

type AuthHandler struct {
	UserStorage *storage.UserStorage
}

func NewAuthHandler(userStorage *storage.UserStorage) *AuthHandler {
	return &AuthHandler{
		UserStorage: userStorage,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input models.RegisterInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if input.Email == "" {
		response.WriteError(w, http.StatusBadRequest, "Email cannot be empty!")
		return
	}

	if input.Password == "" {
		response.WriteError(w, http.StatusBadRequest, "Password cannot be empty!")
		return
	}

	if len(input.Password) < 8 {
		response.WriteError(w, http.StatusBadRequest, "Password must contain at least 8 characters")
		return
	}

	_, err = h.UserStorage.GetUserByEmail(input.Email)
	if err == nil {
		response.WriteError(w, http.StatusConflict, "User with this email already exists!")
		return

	}
	if err != sql.ErrNoRows {
		response.WriteError(w, http.StatusInternalServerError, "Failed to check email!")
		return
	}

	passwordHash, err := auth.HashPassword(input.Password)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "password hashing failed!")
		return
	}

	user := models.User{
		Email:        input.Email,
		PasswordHash: passwordHash,
		Role:         "user",
	}

	createdUser, err := h.UserStorage.CreateUser(user)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to create user!")
		return
	}

	response.WriteJSON(w, http.StatusCreated, createdUser)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input models.LoginInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid JSON!")
		return
	}

	if input.Email == "" {
		response.WriteError(w, http.StatusBadRequest, "Invalid email or password!")
		return
	}

	if input.Password == "" {
		response.WriteError(w, http.StatusBadRequest, "Invalid email or password!")
		return
	}

	user, err := h.UserStorage.GetUserByEmail(input.Email)

	if err != nil {
		response.WriteError(w, http.StatusUnauthorized, "Invalid email or password!")
		return
	}

	if !auth.CheckPassword(input.Password, user.PasswordHash) {
		response.WriteError(w, http.StatusUnauthorized, "Invalid email or password!")
		return
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to generate token!")
		return
	}

	loginResponse := models.LoginResponse{
		Token: token,
	}
	response.WriteJSON(w, http.StatusOK, loginResponse)
}
