package api

type ServerConfiguration struct {
	DatabaseConfiguration *DBConfig `json:"databaseConfiguration" yaml:"databaseConfiguration,omitempty"`
	ListeningAddr         *string   `json:"listeningAddr" yaml:"listeningAddr,omitempty"`
	CacheInSeconds        *int      `json:"cacheInSeconds" yaml:"cacheInSeconds,omitempty"`
}

type DBConfig struct {
	DatabaseType          string `json:"databaseType" yaml:"databaseType,omitempty"`
	SeedDatabaseWhenEmpty bool   `json:"seedDatabaseWhenEmpty" yaml:"seedDatabaseWhenEmpty,omitempty"`
	ConnectionString      string `json:"connectionString" yaml:"connectionString,omitempty"`

	// ConnectionStringSecret represents the configuration for a GCP secret used for the connection string.
	// If this is set, then the connection string is retrieved from the secret and value of ConnectionString is ignored.
	ConnectionStringSecret *GcpSecretConfig `json:"connectionStringSecret" yaml:"connectionStringSecret,omitempty"`
}

type GcpSecretConfig struct {
	ProjectId string `json:"projectId" yaml:"projectId,omitempty"`
	SecretId  string `json:"secretId" yaml:"secretId,omitempty"`
	Revision  string `json:"revision" yaml:"revision,omitempty"`
}

type ctxConfigKey struct{}
