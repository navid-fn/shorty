package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/navid-fn/shorty/internal/api/model"
	"github.com/navid-fn/shorty/internal/store"
	"github.com/navid-fn/shorty/internal/utils"
)

type UserHandler struct {
	userStore store.UserStore
}

func NewUserHandler(userStore store.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}

func (ul *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.UserRegister
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Basic validation
	if req.Username == "" || req.Email == "" || req.Password == "" {
		utils.WriteError(w, http.StatusBadRequest, "Username, email, and password are required")
		return
	}

	// Create the user
	user, err := ul.userStore.CreateUser(&req)
	if err != nil {
		fmt.Println("error Occurred", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, http.StatusOK, user)
}

func (ul *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req model.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Username == "" || req.Password == "" {
		utils.WriteError(w, http.StatusBadRequest, "Username and password are required")
		return
	}

	result, err := ul.userStore.Authenticate(&req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid username or password")
		return
	}

	utils.WriteJson(w, http.StatusOK, result)

}


