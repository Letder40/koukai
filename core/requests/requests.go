package requests

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ServerStrapiRequest(method string, endpoint string, body io.Reader) (string, error) {
    client := &http.Client{}
    url := "http://127.0.0.1:1337" + endpoint

    req, err := http.NewRequest(method, url, body)
    if err != nil {
        return "", err
    }

    auth := fmt.Sprintf("Bearer %s", strings.TrimSuffix(getAppKey(), "\n"))
    req.Header.Add("Authorization", auth)
    req.Header.Add("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }

    if resp.StatusCode > 299 {
        respBody, _ := io.ReadAll(resp.Body)
        return string(respBody), errors.New("Strapi endpoint sent " + resp.Status)
    }     

    respBody, _ := io.ReadAll(resp.Body)
    return string(respBody), nil 
}

func UserStrapiRequest(endpoint string, jwt string) (string, error) {
    client := &http.Client{}
    url := "http://127.0.0.1:1337" + endpoint

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", err
    }

    auth := fmt.Sprintf("Bearer %s", jwt)
    req.Header.Add("Authorization", auth)
    req.Header.Add("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }

    if strings.Split(resp.Status, " ")[0] == "200" {
        respBody, _ := io.ReadAll(resp.Body)
        println("200 OK: api_token validated")
        return string(respBody), nil 
    }     

    return "", errors.New("Strapi endpoint sent " + resp.Status)
}
