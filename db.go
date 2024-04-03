// go:

package main

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GameDb struct {
	ID      int    `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`
	Name    string `gorm:"column:name;type:text(50)" json:"name"`
	Pr      string `gorm:"column:pr;type:text" json:"pr"` //
	Zr      string `gorm:"column:zr;type:text" json:"zr"` //
	Version string `gorm:"column:zr;type:text(10)" json:"version"`
}

var db *gorm.DB

func GetDbData(ctx context.Context) ([]GameDb, error) {
	var err error
	if db == nil {
		db, err = InitDB()
		if err != nil {
			return nil, err
		}
	}
	data := []GameDb{}
	db.Find(&data)
	runtime.EventsEmit(ctx, "db_data", data)
	return data, nil
}

func InitDB() (*gorm.DB, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// dbFile := dir + string(filepath.Separator) + "config" + string(filepath.Separator) + "log.db"
	dbFile := dir + string(filepath.Separator) + "config" + string(filepath.Separator) + "log.db"
	db1, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db = db1
	//
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
