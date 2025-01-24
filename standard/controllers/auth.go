package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/HarshThakur1509/boilerplate/standard/initializers"
	"github.com/markbates/goth/gothic"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	// Get user from the context
	userID := r.Context().Value("userID")

	var user models.User
	initializers.DB.First(&user, "id = ?", userID)

	// Respond with user data
	json.NewEncoder(w).Encode(user)
}

// CUSTOM AUTHENTICATION

func CustomRegister(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string
		Password string
		Name     string
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Failed to Read Body", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusBadRequest)
		return

	}

	user := models.User{Email: body.Email, Password: string(hash), Name: body.Name}

	result := initializers.DB.FirstOrCreate(&user, "email = ?", user.Email)
	if result.Error != nil {
		http.Error(w, "Failed to Create User", http.StatusBadRequest)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Signup successful"})
}

func CustomLogin(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string
		Password string
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Failed to Read body", http.StatusBadRequest)
		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	// Save user ID in the session
	var id string = strconv.FormatUint(uint64(user.ID), 10)
	err = gothic.StoreInSession("user_id", id, r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Return an empty JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Login successful"})
}

func ValidateUser(w http.ResponseWriter, r *http.Request) {
	// Get user from the context
	userID := r.Context().Value("userID")

	// Return an empty JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]interface{}{"userID": userID})
}

// OAUTH AUTHENTICATION
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {

	// Finalize the authentication process
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		log.Println(err)
		return
	}

	// Save user to the database
	userModel := models.User{
		Name:  user.Name,
		Email: user.Email,
	}

	result := initializers.DB.FirstOrCreate(&userModel, "email = ?", userModel.Email)
	if result.Error != nil {
		http.Error(w, "Failed to Create User", http.StatusBadRequest)
		return

	}

	// Save user ID in the session
	var id string = strconv.FormatUint(uint64(userModel.ID), 10)
	err = gothic.StoreInSession("user_id", id, r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Redirect to the secure area
	redirectSecure := os.Getenv("REDIRECT_SECURE")
	if redirectSecure == "" {
		redirectSecure = "http://localhost"
	}

	http.Redirect(w, r, redirectSecure, http.StatusFound)
}

func GothLogout(w http.ResponseWriter, r *http.Request) {
	// Clear session
	err := gothic.Logout(w, r)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	// Redirect to login page
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
