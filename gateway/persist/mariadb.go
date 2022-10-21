package persist

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/gopperin/sme-mini/gateway/config"
	mystore "github.com/gopperin/sme-mini/types/mariadb"
)

// GMariadb GMariadb
var GMariadb Mariadb

// Mariadb Mariadb
type Mariadb struct {
	db *gorm.DB
}

// Init Init
func (maria *Mariadb) Init() error {
	db, err := gorm.Open(config.MariaDB.Dialect, config.MariaDB.URL)
	if err != nil {
		return err
	}

	db.LogMode(false)
	db.DB().SetMaxIdleConns(config.MariaDB.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MariaDB.MaxOpenConns)
	db.DB().SetConnMaxLifetime(10 * time.Minute)

	maria.db = db

	if !db.HasTable(&mystore.GudpUser{}) {
		db.CreateTable(&mystore.GudpUser{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&mystore.GudpUser{})
	}

	if !db.HasTable(&mystore.GudpUserProfile{}) {
		db.CreateTable(&mystore.GudpUserProfile{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&mystore.GudpUserProfile{})
	}

	return nil
}

// Close Close
func Close() error {
	err := GMariadb.db.Close()
	if err != nil {
		fmt.Println("mariadb close error", err.Error())
	}
	fmt.Println("mariadb close")
	return nil
}
