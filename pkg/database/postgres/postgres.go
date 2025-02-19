package postgres

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
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

	fmt.Printf("connectionString: %s", connectionString)
	prxConfig, err := pgx.ParseConfig(connectionString)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	var opts []cloudsqlconn.Option

	opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()))
	opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithDialIAMAuthN(true)))

	d, err := cloudsqlconn.NewDialer(context.Background(), opts...)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// Use the Cloud SQL connector to handle connecting to the instance.
	// This approach does *NOT* require the Cloud SQL proxy.
	prxConfig.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, connectionName)
	}
	dbURI := stdlib.RegisterConnConfig(prxConfig)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        dbURI,
	}))

	if err != nil {
		return fmt.Errorf("failed to connect database - %v", err)
	}

	v.DB = db
	return nil
	// db, err := gorm.Open(postgres.Open(connectionString), config)
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to connect database - %v", err))
	// }
	// v.DB = db

}

// GetDB returns the underlying *gorm.DB instance of the Database.
func (v *Database) GetDB() *gorm.DB {
	return v.DB
}
