package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/parsaaes/another-url-shortener/config"
	"time"
)

const connectTimeout = 60 * time.Second

func CreatePostgresDB() (*gorm.DB, error) {
	cfg := config.Cfg.Postgres
	url := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s connect_timeout=%d sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.DB, cfg.Pass, connectTimeout)
	postgresDB, err := gorm.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return postgresDB, nil
}
