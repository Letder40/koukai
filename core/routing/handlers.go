package routing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"koukai/requests"
	"koukai/utils"
	"net/http"
	"sync"
)

type userLoginData struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type userSignupData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserData struct {
	ID         int    `json:"id"`
	DocumentId string `json:"documentId"`
	Username   string `json:"username"`
	Email      string `json:"email"`
}

type AppRuntime struct {
	UserData map[string]UserData
	mutex    sync.RWMutex
}

func (appRuntime *AppRuntime) initRuntime() {
	appRuntime.UserData = make(map[string]UserData)
	appRuntime.mutex = sync.RWMutex{}
}

func (app *AppRuntime) cacheUser(jwt string, UserData UserData) {
	app.mutex.Lock()
	app.UserData[jwt] = UserData
	app.mutex.Unlock()
}

func (app *AppRuntime) userHandler(w http.ResponseWriter, r *http.Request) {
	jwt, err := requests.GetUserJWT(r)
	if err != nil {
		http.Error(w, "Forbidden: Invalid or missing JWT", http.StatusForbidden)
		return
	}

	app.mutex.RLock()
	userData, exists := app.UserData[jwt]
	app.mutex.RUnlock()

	if !exists {
		resp, err := requests.UserStrapiRequest("/api/users/me", jwt)
		if err != nil {
			println("Strapi request error:", err.Error())
			http.Error(w, fmt.Sprintf("Failed to fetch user data: %v", err), http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal([]byte(resp), &userData); err != nil {
			println("Unmarshal error:", err.Error())
			println("Response body:", string(resp))
			http.Error(w, "Unable to parse user data", http.StatusInternalServerError)
			return
		}

		app.cacheUser(jwt, userData)
	}

	if err := utils.SendJsonFrom(w, userData); err != nil {
		println("SendJson error:", err.Error())
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}

	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var UserData userLoginData
	err = json.Unmarshal(body, &UserData)
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

func handleSignup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var userSignupData userSignupData
	err = json.Unmarshal(body, &userSignupData)
	if err != nil {
		http.Error(w, "Invalid Json format", http.StatusBadRequest)
		return
	}

	if resp, err := requests.ServerStrapiRequest("POST", "/api/users", bytes.NewReader(body)); err != nil {
		http.Error(w, fmt.Sprintf("Strapi error: %s", resp), http.StatusBadRequest)
		return
	}
}
