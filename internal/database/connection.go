package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type connConfig struct {
	Host          string
	Port          string
	DBName        string
	Username      string
	Password      string
	SslMode       string
	SslRootCert   string
	SslClientCert string
	SslClientKey  string
	LogMode       bool
	LogQueries    bool
	Timeout       string
}

func connectDatabase() (*sqlx.DB, error) {
	config := getConfig()
	db, err := sqlx.Connect("postgres", connString(config))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDatabaseConnection() (*sqlx.DB, error) {
	var err error
	var db *sqlx.DB

	db, err = connectDatabase()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connString(c *connConfig) string {
	template := "host=%v port=%v user=%v dbname=%v password=%v sslmode=%v sslrootcert=%v sslcert=%v sslkey=%v connect_timeout=%v"
	return fmt.Sprintf(template, c.Host, c.Port, c.Username, c.DBName, c.Password, c.SslMode, c.SslRootCert, c.SslClientCert, c.SslClientKey, c.Timeout)
}

func getConfig() *connConfig {
	config := &connConfig{
		Host:          getEnv("DB_HOST", "localhost"),
		Port:          getEnv("DB_PORT", "5432"),
		DBName:        getEnv("DB_NAME", "usuarios-db"),
		Username:      getEnv("DB_USERNAME", "postgres"),
		Password:      getEnv("DB_PASSWORD", "postgres"),
		SslMode:       getEnv("DB_SSLMODE", "disable"),
		SslRootCert:   getEnv("DB_SSL_ROOT_CERT", "false"),
		SslClientCert: getEnv("DB_SSL_CLIENT_CERT", "false"),
		SslClientKey:  getEnv("DB_SSL_CLIENT_KEY", "false"),
		LogMode:       getEnvAsBool("DB_LOGMODE", true),
		LogQueries:    getEnvAsBool("DB_LOGQUERIES", true),
		Timeout:       getEnv("DB_TIMEOUT", "2"),
	}

	return config
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	if defaultValue != "" {
		return defaultValue
	}
	panic(fmt.Sprintf("Environment variable %v has not value.", key))
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if strValue, exists := os.LookupEnv(key); exists {
		if value, err := strconv.ParseBool(strValue); err == nil {
			return value
		}
	}

	return defaultValue
}
