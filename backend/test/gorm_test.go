package test

import (
	"fmt"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGormTest(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	// db.Create()
	fmt.Println(db)

}
