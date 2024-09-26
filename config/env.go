package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// DSN - Data source name
func GetDSN() string {
	godotenv.Load()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	print(dsn)

	return dsn
}
