package structs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/xristoskrik/yourmusic/internal/database"
)

func TestUserCreateHandler(t *testing.T) {
	url := "http://localhost:8080/api/users"
	type credentials struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	creds := credentials{
		Password: "test",
		Email:    "test6@gmail.com",
	}
	jsonData, err := json.Marshal(creds)
	if err != nil {
		t.Errorf("failed to Marshal")
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf("failed to post")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusInternalServerError {
		t.Log("User already created")
		t.Skip()
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusCreated)
	}

	var res database.User
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&res)
	if err != nil {
		t.Errorf("failed to decode")
	}
	fmt.Printf("%+v\n", res)

}
