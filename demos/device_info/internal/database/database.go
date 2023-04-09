package database

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"git.sofunny.io/data-analysis/device_info/internal/config"
)

type DB struct {
	*gorm.DB
}

func New(c *config.DatabaseCfg) (*DB, error) {
	var (
		gormConf = &gorm.Config{
			Logger: &Logger{},
		}
		db  *gorm.DB
		err error

		opts = c.Opts
	)

	if opts == nil {
		opts = make(map[string]string)
	}

	switch c.Type {
	case "postgres":
		port := c.Port
		if port == 0 {
			port = 5432
		}
		dsnB := []string{
			fmt.Sprintf("dbname=%s", c.DBName),
			fmt.Sprintf("host=%s", c.Host),
			fmt.Sprintf("port=%d", port),
			fmt.Sprintf("user=%s", c.User),
			fmt.Sprintf("password=%s", c.Password),
		}

		for k, v := range opts {
			dsnB = append(dsnB, fmt.Sprintf("%s=%q", k, v))
		}
		dsn := strings.Join(dsnB, " ")
		db, err = gorm.Open(postgres.Open(dsn), gormConf)
	default:
		return nil, fmt.Errorf("unsupported database type %s", c.Type)
	}
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxIdleTime(time.Duration(c.Pool.ConnMaxIdleTime) * time.Second)
	sqlDB.SetConnMaxLifetime(time.Duration(c.Pool.ConnMaxLifeTime) * time.Second)
	sqlDB.SetMaxIdleConns(c.Pool.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.Pool.MaxOpenConns)

	logger := log.Logger.With().Str("m", "gorm").Logger()
	db = db.WithContext(logger.WithContext(context.Background()))

	return &DB{db}, err
}

func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		return err
	}
	return nil
}

func encodeDBOpts(m map[string]string) string {
	values := make(url.Values)
	for k, v := range m {
		values.Set(k, v)
	}
	return values.Encode()
}
