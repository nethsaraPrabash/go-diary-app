package test

import (
	"encoding/json"
	"net/http"
	"testing"
	"bytes"

)

func TestRegister(t *testing.T) {

	user := map[string]string{
		"username": "test",
		"password": "test123",
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post("http://localhost:5050/auth/register", "application/json", bytes.NewBuffer(jsonUser))

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200; got %v", resp.StatusCode)
	}
}

func TestLogin(t *testing.T) {
	
	user := map[string]string{
		"username": "test",
		"password": "test123",
	}

	jsonUser, err := json.Marshal(user)

	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post("http://localhost:5050/auth/login", "application/json", bytes.NewBuffer(jsonUser))

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200; got %v", resp.StatusCode)
	}

}