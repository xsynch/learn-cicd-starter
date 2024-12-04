package auth

import (
	"net/http"

	"testing"
)

func TestAuth(t *testing.T) {

	req, _ := http.NewRequest("GET", "http.example.com", nil)

	req.Header.Set("Authorization", "ApiKey thisisatest")

	expected := "thisisatest"
	val, _ := GetAPIKey(req.Header)

	if expected != val {
		t.Errorf("unexpected reuturn: got %v want %v", val, expected)
	}

}

func TestAuthFailure(t *testing.T) {

	req, _ := http.NewRequest("GET", "http.example.com", nil)

	req.Header.Set("NoAuthorization", "ApiKeythisisatest")

	expected := ErrNoAuthHeaderIncluded
	val, err := GetAPIKey(req.Header)

	if expected != err {
		t.Errorf("unexpected reuturn: got %v want %v", val, expected)
	}

}
