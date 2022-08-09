package accounts

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"yatter-backend-go/app/config"
	"yatter-backend-go/app/dao"
	"yatter-backend-go/app/domain/object"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		requestBody    io.Reader
		wantStatusCode int
		wantUsername   string
	}{
		{
			name:           "example",
			url:            "http://localhost:1234/",
			wantUsername:   "username",
			requestBody:    strings.NewReader(`{"username": "username", "password": "password"}`),
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "empty_username",
			url:            "http://localhost:1234/",
			requestBody:    strings.NewReader(`{"username": "", "password": "password"}`),
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "empty_password",
			url:            "http://localhost:1234/",
			requestBody:    strings.NewReader(`{"username": "username", "password": ""}`),
			wantStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := dao.New(config.MySQLConfig())
			if err != nil {
				t.Fatal(err)
			}
			err = d.InitAll()
			if err != nil {
				t.Fatal(err)
			}
			res, err := http.Post(tt.url, "application/json", tt.requestBody)
			if err != nil {
				t.Fatal(err)
			}
			gotStatusCode := res.StatusCode
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("unexpected status code %s got = %v, want %v", tt.url, gotStatusCode, tt.wantStatusCode)
			}
			if gotStatusCode != http.StatusOK {
				return
			}
			var gotBody object.Account
			if err := json.NewDecoder(res.Body).Decode(&gotBody); err != nil {
				t.Fatal(err)
			}

			if err != nil {
				t.Fatal(err)
			}
			if gotBody.Username != tt.wantUsername {
				t.Errorf("unexpected body %s got = %v, want %v", tt.url, gotBody.Username, tt.wantUsername)
			}
		})
	}
}
