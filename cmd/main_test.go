package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_api(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/health", GetHealth)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	var have map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &have); err != nil {
		t.Errorf("Could not unmarshal response: %v", err)
	}

	want := "Hello"
	if have["message"] != want {
		t.Errorf("handler returned unexpected body: got %v want %v", have["message"], want)
	}
}
