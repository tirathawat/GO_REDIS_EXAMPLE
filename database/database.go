package database

import (
	"example/models/storages"
	"fmt"
	"math/rand"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

type IDatabase interface {
	GetDB() *gorm.DB
	CloseDB() error
	OpenConnection() error
	FeedData() error
}

var instantiated *database = nil

func New() *database {
	if instantiated == nil {
		instantiated = new(database)
	}
	return instantiated
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}

func (db *database) OpenConnection() error {
	var err error
	dsn := getDSN()
	sql := mysql.Open(dsn)
	db.db, err = gorm.Open(sql)
	return err
}

func (db *database) GetDB() *gorm.DB {
	return db.db
}

func (db *database) CloseDB() error {
	sql, err := db.db.DB()
	if err != nil {
		return err
	}
	err = sql.Close()
	return err
}

func (db *database) FeedData() error {
	var count int64
	db.db.Model(&storages.DemoDB{}).Count(&count)
	if count > 0 {
		return nil
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	products := []storages.DemoDB{}
	for i := 0; i < 5000; i++ {
		products = append(products, storages.DemoDB{
			Name:     fmt.Sprintf("Demo %v", i+1),
			Priority: random.Intn(100),
		})
	}
	return db.db.Create(&products).Error
}
