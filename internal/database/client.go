package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		"localhost",
		5431,
		"postgres",
		"my-postgresql",
		"postgres",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy:         schema.NamingStrategy{TablePrefix: "wisdom."},
		FullSaveAssociations:   false,
		Logger:                 nil,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		DB: db,
	}

	return client, nil
}

func (c Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as 'ready'").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready != "1" {
		return true
	}
	return false
}
