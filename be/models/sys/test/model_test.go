package test

import (
	"testing"

	_ "gapp.models.sys"
	"github.com/vn-go/dx"
)

var dsn string = "postgres://postgres:123456@localhost:5432/gapp?sslmode=disable"

func TestModel(t *testing.T) {
	db, err := dx.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

}
