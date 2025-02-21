package requests

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

const token = "Bearer b9ca40a9d8400d16e23a6897b1e903bc2be9238e6d61df5e30e3c3593aaef8f570a56e5c11533ef7af9a49cbeeb858ddbe3fe3ef73dbdf2fbc6d936f3757d88aac3c2edb93271e9fd606c124843aba6e84079f86b9b6e26921bbc47435ba4e1d1ff978334cf65b4f64d24b33721cb2b9f3998424baae813d93094d70b4aaf671"

func RequestStrapi(endpoint string) (string, error) {
    client := &http.Client{}
    url := "http://127.0.0.1:1337/api/" + endpoint
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", err
    }
    req.Header.Add("Authorization", token)
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }

    if strings.Split(resp.Status, " ")[0] == "200" {
        respBody, _ := io.ReadAll(resp.Body)
        return string(respBody), nil 
    }     

    return "", errors.New("Strapi endpoint sent " + resp.Status)
}
