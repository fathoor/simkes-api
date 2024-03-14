package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewDB(cfg *Config, log *zerolog.Logger) *gorm.DB {
	var (
		host         = cfg.Get("PG_HOST")
		port         = cfg.Get("PG_PORT")
		user         = cfg.Get("PG_USER")
		password     = cfg.Get("PG_PASSWORD")
		dbname       = cfg.Get("PG_NAME")
		connOpen     = cfg.GetInt("PG_CONN_OPEN")
		connIdle     = cfg.GetInt("PG_CONN_IDLE")
		connLifeTime = cfg.GetInt("PG_CONN_LIFETIME")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, user, password, dbname)
	dialect := postgres.Open(dsn)

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	psql, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create database instance")
	}

	psql.SetMaxOpenConns(connOpen)
	psql.SetMaxIdleConns(connIdle)
	psql.SetConnMaxLifetime(time.Minute * time.Duration(connLifeTime))

	return db
}
