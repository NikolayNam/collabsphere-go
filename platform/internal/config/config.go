package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	TZ  string `env:"TZ" envDefault:"UTC"`
	APP App
	DB  DB
}

type App struct {
	Title        string        `env:"APPLICATION_TITLE,required"`
	Version      string        `env:"APPLICATION_VERSION,required"`
	Address      string        `env:"APPLICATION_ADDRESS" envDefault:"0.0.0.0:8080"`
	TimeoutRead  time.Duration `env:"APPLICATION_TIMEOUT_READ" envDefault:"15s"`
	TimeoutWrite time.Duration `env:"APPLICATION_TIMEOUT_WRITE" envDefault:"15s"`
	TimeoutIdle  time.Duration `env:"APPLICATION_TIMEOUT_IDLE" envDefault:"60s"`
	Debug        bool          `env:"APPLICATION_DEBUG" envDefault:"false"`
}

type DB struct {
	Host         string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port         int    `env:"POSTGRES_PORT" envDefault:"5432"`
	DBName       string `env:"POSTGRES_DB" envDefault:"postgres"`
	DBSchema     string `env:"POSTGRES_SCHEMA,required"`
	Username     string `env:"POSTGRES_USER" envDefault:"postgres"`
	Password     string `env:"POSTGRES_PASSWORD"`
	PasswordFile string `env:"POSTGRES_PASSWORD_FILE"`
	Debug        bool   `env:"POSTGRES_DEBUG" envDefault:"false"`
}

func New() *Config {
	var c Config

	if err := env.Parse(&c); err != nil {
		log.Fatalf("failed to parse env: %s", err)
	}

	if err := c.DB.loadSecret(); err != nil {
		log.Fatalf("failed to load DB secret: %s", err)
	}

	if err := applyTZ(c.TZ); err != nil {
		log.Fatalf("invalid TZ: %s", err)
	}

	return &c
}

// --- Secret Loader ---

func (db *DB) loadSecret() error {
	if db.Password != "" {
		return nil
	}

	if db.PasswordFile != "" {
		b, err := os.ReadFile(db.PasswordFile)
		if err != nil {
			return err
		}
		db.Password = strings.TrimSpace(string(b))
	}

	return nil
}

// --- Timezone Apply ---

func applyTZ(tz string) error {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return err
	}
	time.Local = loc
	return nil
}

// --- DSN Builder (не логируй его целиком) ---

func (db *DB) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable&search_path=%s",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.DBName,
		db.DBSchema,
	)
}
