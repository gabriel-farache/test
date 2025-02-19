package database

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DBInterface interface {
	SetLogger(*zap.SugaredLogger)
	Init(string, string, *gorm.Config) error
	GetDB() *gorm.DB
}

// ConnectionStringToDSN takes a connection string and returns the DSN (Data Source Name) and connection name.
// The connection string should be in the format of space-separated key=value pairs.
// It extracts the connection name, database, user, and password from the connection string.
// If any required parameter is missing, it returns an error.
// If the password is missing, it returns the DSN without the password.
// Parameters:
//   - connectionString: The connection string to parse.
//
// Returns:
//   - string: The DSN (Data Source Name) for the connection.
//   - string: The connection name.
//   - error: An error if any required parameter is missing.
func ConnectionStringToDSN(connectionString string) (string, string, error) {
	if !strings.Contains(connectionString, " ") {
		return "", "", fmt.Errorf("invalid connection string. Please supply a valid connection string with space separated key=value pairs")
	}
	cstrArr := strings.Split(connectionString, " ")

	connectionName := ""
	db := ""
	user := ""
	password := ""
	for _, paramAndVal := range cstrArr {
		paramValArr := strings.SplitN(paramAndVal, "=", 2)
		switch paramValArr[0] {
		case "connectionname":
			connectionName = paramValArr[1]
		case "database":
			db = paramValArr[1]
		case "user":
			user = paramValArr[1]
		case "password":
			password = paramValArr[1]
		}
	}

	if connectionName == "" {
		return "", "", fmt.Errorf("connection string is missing the connectionname parameter")
	}
	if db == "" {
		return "", "", fmt.Errorf("connection string is missing the database parameter")
	}
	if user == "" {
		return "", "", fmt.Errorf("connection string is missing the user parameter")
	}

	if password == "" {
		return fmt.Sprintf("user=%s database=%s", user, db), connectionName, nil
	} else {
		return fmt.Sprintf("user=%s password=%s database=%s", user, password, db), connectionName, nil
	}

}
