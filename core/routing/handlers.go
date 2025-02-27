package routing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"koukai/requests"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    body, err := io.ReadAll(r.Body);
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        return
    }

    var userData UserLoginData;
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

func HandleSingup(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close();

    body, err := io.ReadAll(r.Body);
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        return
    }

    var userData UserData;
    err = json.Unmarshal(body, &userData)
    if err != nil {
        http.Error(w, "Invalid Json format", http.StatusBadRequest)
        return
    }

    if _, err := requests.ServerStrapiRequest("POST", "/api/users", bytes.NewReader(body)); err != nil {
        http.Error(w, "Strapi error", http.StatusBadRequest)
        return
    }
}
