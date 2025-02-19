package api

import (
	"context"
	"fmt"

	"github.com/IaC/go-kcloutie/pkg/database"
	"github.com/IaC/go-kcloutie/pkg/database/postgres"
	"github.com/IaC/go-kcloutie/pkg/database/sqlite"
	"github.com/IaC/go-kcloutie/pkg/gcp"
	"github.com/IaC/go-kcloutie/pkg/model"
	"gorm.io/gorm"
)

func (c *ServerConfiguration) configureDB(ctx context.Context) (database.DBInterface, error) {
	connectionName := ""
	err := c.CheckDbServerConfig()
	if err != nil {
		return nil, err
	}

	if c.DatabaseConfiguration == nil {
		return nil, nil
	}

	if c.DatabaseConfiguration.ConnectionStringSecret != nil {
		secClient := gcp.FromCtx(ctx)
		if secClient != nil {
			defer secClient.Close()
		}
		secretData, err := gcp.GetSecret(ctx, secClient, c.DatabaseConfiguration.ConnectionStringSecret.ProjectId, c.DatabaseConfiguration.ConnectionStringSecret.SecretId, c.DatabaseConfiguration.ConnectionStringSecret.Revision)
		if err != nil {
			return nil, fmt.Errorf("failed to get connection string secret data - %v", err)
		}
		c.DatabaseConfiguration.ConnectionString, connectionName, err = database.ConnectionStringToDSN(string(secretData))
		if err != nil {
			return nil, fmt.Errorf("failed to parse connection string - %v", err)
		}
	} else {
		c.DatabaseConfiguration.ConnectionString, connectionName, err = database.ConnectionStringToDSN(c.DatabaseConfiguration.ConnectionString)
		if err != nil {
			return nil, fmt.Errorf("failed to parse connection string - %v", err)
		}
	}

	var dbInterface database.DBInterface

	if c.DatabaseConfiguration.DatabaseType == "sqlite" {
		dbInterface, err = c.initializeSqlite(connectionName)
		if err != nil {
			return nil, err
		}
	}

	if c.DatabaseConfiguration.DatabaseType == "postgres" {

		dbInterface, err = c.initializePostgres(connectionName)
		if err != nil {
			return nil, err
		}
	}

	err = dbInterface.GetDB().AutoMigrate(&model.Widget{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database - %v", err)
	}
	if !c.DatabaseConfiguration.SeedDatabaseWhenEmpty {
		return dbInterface, nil
	}

	var widgets []model.Widget
	if result := dbInterface.GetDB().Find(&widgets); result.Error != nil {
		return dbInterface, fmt.Errorf("failed to query widgets - %v", result.Error)
	}

	c.seedDb(widgets, dbInterface)

	return dbInterface, nil

}

func (c *ServerConfiguration) initializeSqlite(connectionName string) (database.DBInterface, error) {
	db := sqlite.New()
	err := db.Init(c.DatabaseConfiguration.ConnectionString, "", &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to configure the sqlite database - %v", err)
	}
	return db, nil
}

func (c *ServerConfiguration) initializePostgres(connectionName string) (database.DBInterface, error) {

	if connectionName == "" {
		return nil, fmt.Errorf("connection name is required for postgres database")
	}
	db := postgres.New()
	err := db.Init(c.DatabaseConfiguration.ConnectionString, connectionName, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to configure the postgres database - %v", err)
	}
	return db, nil
}

func (*ServerConfiguration) seedDb(widgets []model.Widget, dbInterface database.DBInterface) {
	if len(widgets) == 0 {
		newWidgets := model.GetSeedData(20)
		for _, widget := range newWidgets {
			dbInterface.GetDB().Create(&widget)
		}
	}
}

func (c *ServerConfiguration) CheckDbServerConfig() error {
	if c.DatabaseConfiguration == nil || c.DatabaseConfiguration.DatabaseType == "" {
		fmt.Println("No database type specified. Skipping database configuration")
		return nil
	}
	if c.DatabaseConfiguration.DatabaseType != "postgres" && c.DatabaseConfiguration.DatabaseType != "sqlite" {
		return fmt.Errorf("invalid value of '%s' for dbtype. You must supply 'sqlite' or 'postgres' for the dbtype switch value", c.DatabaseConfiguration.DatabaseType)
	}
	return nil
}

func (c *ServerConfiguration) DBEnabled() bool {
	if c.DatabaseConfiguration == nil {
		return false
	}
	return c.DatabaseConfiguration.DatabaseType != ""
}
