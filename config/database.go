package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/prolifel/kaching/models"
	"golang.org/x/exp/slices"
)

type DB struct {
	*sqlx.DB
}

func NewDatabase(driver string, connectionString string, connLifeTime int64) (db *DB, err error) {
	if !slices.Contains(models.AppAvaialbleDatabase, driver) {
		err = errors.New("driver not available")
		return
	}

	dbx, err := sqlx.Connect(driver, connectionString)
	if err != nil {
		err = errors.New(fmt.Sprintf("[DatabaseConnection] failed connect to database %+v", err))
		return
	}

	connectionTTL := os.Getenv(models.EnvDBConnLifetime)
	if connectionTTL == "" {
		connectionTTL = "15"
	}

	connectionTTLInt, err := strconv.Atoi(connectionTTL)
	if err != nil {
		err = errors.New(fmt.Sprintf("[DatabaseConnection] while convert string to int %+v", err))
		return
	}

	connectionMaxIdle := os.Getenv(models.EnvDBConnMaxIdle)
	if connectionMaxIdle == "" {
		connectionMaxIdle = "5"
	}

	connectionMaxIdleInt, err := strconv.Atoi(connectionMaxIdle)
	if err != nil {
		err = errors.New(fmt.Sprintf("[DatabaseConnection] while convert string to int %+v", err))
		return
	}

	connectionMaxOpen := os.Getenv(models.EnvDBConnMaxOpen)
	if connectionMaxOpen == "" {
		connectionMaxOpen = "0"
	}

	connectionMaxOpenInt, err := strconv.Atoi(connectionMaxOpen)
	if err != nil {
		err = errors.New(fmt.Sprintf("[DatabaseConnection] while convert string to int %+v", err))
		return
	}

	dbx.SetConnMaxLifetime(time.Minute * time.Duration(connectionTTLInt))
	dbx.SetMaxIdleConns(connectionMaxIdleInt)
	dbx.SetMaxOpenConns(connectionMaxOpenInt)

	db = &DB{
		dbx,
	}

	return
}

func NewPostgresqlConnection(connectionString string, connLifeTime int64) (*DB, error) {
	return NewDatabase(models.ConstDatabaseTypePostgresql, connectionString, connLifeTime)
}
