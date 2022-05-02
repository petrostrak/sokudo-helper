package data

import (
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNew(t *testing.T) {
	fakeDB, _, _ := sqlmock.New()
	defer fakeDB.Close()

	_ = os.Setenv("DATABASE_TYPE", "postgres")
	m := New(fakeDB)

	if fmt.Sprintf("%T", m) != "data.Models" {
		t.Error("wrong type", fmt.Sprintf("%T", m))
	}

	_ = os.Setenv("DATABASE_TYPE", "mysql")
	m = New(fakeDB)

	if fmt.Sprintf("%T", m) != "data.Models" {
		t.Error("wrong type", fmt.Sprintf("%T", m))
	}
}
