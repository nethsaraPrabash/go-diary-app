package test

import (
	"encoding/json"
	"net/http"
	"testing"
	"bytes"

)

const authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE3MzgzNDUxMjQsImlhdCI6MTczODM0MzEyNCwiaWQiOjEwfQ.ldMrg6XLZfO2E4qVPjmvxVbEi6R65xa6vSmFdhIfaDA"

func TestAddEntry(t *testing.T) {

	entry := map[string]string{
		"content": "Test content for the diary",
	}
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:5050/api/entry", bytes.NewBuffer(jsonEntry))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated { 
		t.Fatalf("Expected status 201; got %v", resp.StatusCode)
	}
}