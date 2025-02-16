package main

import (
	"bytes"
	"encoding/json"
	"go/web-api/internal/auth"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.com",
		Password: "password",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse

	err = json.Unmarshal(body, &resData)

	if err != nil {
		t.Fatal(err)
	}

	if resData.Token == "" {
		t.Fatal("Token is empty")
	}

}

func TestLoginFail(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.com",
		Password: "passwords",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, res.StatusCode)
	}
}
