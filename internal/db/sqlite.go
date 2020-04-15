package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kumato/kumato/internal/types"
	"log"
	"os"
)

var db *gorm.DB

func Connect(data string, debug bool) {
	var err error

	if data == "" {
		data = os.TempDir()
	}

	db, err = gorm.Open("sqlite3", data+"/.kumato.db")
	if err != nil {
		log.Fatalln(err)
	}

	db.LogMode(debug)

	if m := os.Getenv("KUMATO_DEV_MODE"); m == "1" || m == "on" {
		db.LogMode(true)
	}

	// init table if db file is brand new
	if !db.HasTable(&types.Task{}) {
		db.CreateTable(&types.Task{})
	}
}
