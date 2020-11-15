package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendSMS(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/sendsms?phone=+6285123123&text=Testing+SMS", nil)
	w := httptest.NewRecorder()

	appRoute().ServeHTTP(w, r)

	var result map[string]interface{}
	json.NewDecoder(w.Body).Decode(&result)
	result_vendor := result["result"].(map[string]interface{})

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, result_vendor["status"], true)
}
