package sqlite

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Logger *zap.SugaredLogger
	DB     *gorm.DB
}

// New creates a new instance of the Database struct.
func New() *Database {
	return &Database{}
}

// SetLogger sets the logger for the Database instance.
// The logger parameter should be a pointer to a zap.SugaredLogger.
// This logger will be used to log messages and errors related to the Database operations.
func (v *Database) SetLogger(logger *zap.SugaredLogger) {
	v.Logger = logger
}

// Init initializes the database connection using the provided connection string, connection name, and GORM configuration.
// It returns an error if initialization or the connection fails.
func (v *Database) Init(connectionString, connectionName string, config *gorm.Config) error {

	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database - %v", err)
	}
	v.DB = db
	return nil
}

// GetDB returns the underlying *gorm.DB instance of the Database.
func (v *Database) GetDB() *gorm.DB {
	return v.DB
}
