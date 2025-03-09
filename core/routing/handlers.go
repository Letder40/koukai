package routing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"koukai/requests"
	"net/http"
	"sync"
)

type userLoginData struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type userSingupData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type userData struct {
	Id         uint64 `json:"id"`
	DocumentId string `json:"documentId"`
	Username   string `json:"username"`
	Email      string `json:"email"`
}

type AppRuntime struct {
	userData map[string]userData
	mutex    sync.RWMutex
}

func sendJsonFrom[T any](w http.ResponseWriter, data T) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, jsonData)
	return nil
}

func (app *AppRuntime) cacheUser(jwt string, userData userData) {
	app.mutex.Lock()
	defer app.mutex.Unlock()
	app.userData[jwt] = userData
}

func (app *AppRuntime) userHandler(w http.ResponseWriter, r *http.Request) {
	app.mutex.RLock()
	defer app.mutex.RUnlock()
	jwt, err := requests.GetUserJWT(r)
	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	userData, exists := app.userData[jwt]
	if !exists {
		resp, err := requests.UserStrapiRequest("/api/users/me", jwt)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal([]byte(resp), &userData)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		app.cacheUser(jwt, userData)
	}

	err = sendJsonFrom(w, userData)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var userData userLoginData
	err = json.Unmarshal(body, &userData)
	if err != nil {
		http.Error(w, "Invalid Json format", http.StatusBadRequest)
		return
	}

	var response string
	if response, err = requests.ServerStrapiRequest("POST", "/api/auth/local", bytes.NewReader(body)); err != nil {
		println(err.Error())
		http.Error(w, "Strapi error", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, response)
}

func handleSingup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var userSingupData userSingupData
	err = json.Unmarshal(body, &userSingupData)
	if err != nil {
		http.Error(w, "Invalid Json format", http.StatusBadRequest)
		return
	}

	if resp, err := requests.ServerStrapiRequest("POST", "/api/users", bytes.NewReader(body)); err != nil {
		http.Error(w, fmt.Sprintf("Strapi error: %s", resp), http.StatusBadRequest)
		return
	}
}
