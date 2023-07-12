package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/prolifel/kaching/models"
)

type Config struct {
	DB      *DB
	RestAPI *fiber.App
}

func New() Config {
	var cfg Config

	InitEnv()
	InitTimezone()

	return cfg
}

func Catch(err error) {
	if err != nil {
		panic(err)
	}
}

func InitEnv() (err error) {
	if err = godotenv.Load(); err != nil {
		return
	}

	return
}

func InitTimezone() (err error) {
	loc := os.Getenv(models.EnvAppTimezone)
	local, err := time.LoadLocation(loc)
	if err != nil {
		return
	}
	time.Local = local
	return
}

func (c *Config) InitProgresql() (err error) {
	var connectionString string

	dbHost := os.Getenv(models.EnvDBHost)
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbUser := os.Getenv(models.EnvDBUser)
	if dbUser == "" {
		dbUser = "user"
	}

	dbPwd := os.Getenv(models.EnvDBPwd)
	if dbPwd == "" {
		dbPwd = ""
	}

	dbPort := os.Getenv(models.EnvDBPort)
	if dbPort == "" {
		dbPort = "5432"
	}

	dbName := os.Getenv(models.EnvDBName)
	if dbName == "" {
		dbName = "kaching"
	}

	dbSSLMode := os.Getenv(models.EnvDBSSLMode)
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}

	connectionString = fmt.Sprintf(`
		host=%s
		port=%s
		user=%s
		password=%s
		dbname=%s
		sslmode=%s
		`, dbHost, dbPort, dbUser, dbPwd, dbName, dbSSLMode)

	db, err := NewPostgresqlConnection(connectionString, 15)
	if err != nil {
		err = errors.New(fmt.Sprintf("[DatabaseConnection] while connecting database connection: %v", err.Error()))
		return
	}
	c.DB = db

	return
}
