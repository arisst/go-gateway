package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Request(url string) map[string]interface{} {

	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)

	if err != nil {
		return nil
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result
}
