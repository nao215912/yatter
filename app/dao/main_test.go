package dao

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"testing"
	"yatter-backend-go/app/config"
)

var db *sqlx.DB

func NewTestDao(t *testing.T, queries ...string) Dao {
	d := &dao{db: db}
	err := d.InitAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, query := range queries {
		_, err = db.Exec(query)
		if err != nil {
			t.Fatal(err)
		}
	}
	return d
}

func TestMain(m *testing.M) {

	d, err := initDb(config.MySQLConfig())
	if err != nil {
		log.Fatalln(err)
	}
	defer d.Close()
	db = d
	os.Exit(m.Run())
}
