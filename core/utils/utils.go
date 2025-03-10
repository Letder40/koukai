package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJsonFrom[T any](w http.ResponseWriter, data T) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))
	return nil
}
