package accounts

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
	"yatter-backend-go/app/app"
	"yatter-backend-go/app/config"
)

func TestMain(m *testing.M) {
	go func() {
		os.Setenv("PORT", "1234")
		a, err := app.NewApp()
		if err != nil {
			log.Fatal(err)
		}
		addr := ":" + strconv.Itoa(config.Port())
		http.ListenAndServe(addr, NewRouter(a))
	}()
	for {
		_, err := http.Get("http://localhost:1234")
		if err == nil {
			break
		}
	}
	os.Exit(m.Run())
}
