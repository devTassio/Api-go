package config

import (
	"os"

	"github.com/devTassio/Api-go/schemas"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/database_file.db"

	// Verifica se o arquivo do banco de dados existe
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")

		// Cria o arquivo e o diretório do banco de dados
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create database directory")
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create database file")
		}
		file.Close()
	}

	// Cria o DB e conecta
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to initialize SQLite: %v", err)
		return nil, err
	}

	logger.Info("SQLite initialized successfully")

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("SQLite auto migration error: %v", err)
		return nil, err
	}

	return db, nil
}
