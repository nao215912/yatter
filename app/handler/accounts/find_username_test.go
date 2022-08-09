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

func TestFindUsername(t *testing.T) {
	tests := []struct {
		name              string
		createUrl         string
		createRequestBody io.Reader
		findUrl           string
		wantUsername      string
		wantStatusCode    int
	}{
		{
			name:              "example",
			createUrl:         "http://localhost:1234/",
			createRequestBody: strings.NewReader(`{"username": "username", "password": "password"}`),
			findUrl:           "http://localhost:1234/username",
			wantUsername:      "username",
			wantStatusCode:    http.StatusOK,
		},
		{
			name:              "no_such_user",
			createUrl:         "http://localhost:1234/",
			createRequestBody: strings.NewReader(`{"username": "username", "password": "password"}`),
			findUrl:           "http://localhost:1234/no_such_user",
			wantStatusCode:    http.StatusBadRequest,
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
			res, err := http.Post(tt.createUrl, "application/json", tt.createRequestBody)
			if err != nil {
				t.Fatal(err)
			}
			if res.StatusCode != http.StatusOK {
				t.Fatalf("%s must return status ok", tt.createUrl)
			}
			res, err = http.Get(tt.findUrl)
			if err != nil {
				t.Fatal(err)
			}
			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("unexpected status code  got = %v, want %v", res.StatusCode, tt.wantStatusCode)
			}
			if res.StatusCode != http.StatusOK {
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
				t.Errorf("unexpected body  got = %v, want %v", gotBody.Username, tt.wantUsername)
			}
		})
	}
}
