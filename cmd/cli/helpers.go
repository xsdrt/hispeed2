package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func setup() {
	err := godotenv.Load()
	if err != nil {
		exitGracefully(err)
	}

	path, err := os.Getwd()
	if err != nil {
		exitGracefully(err)
	}

	his.RootPath = path
	his.DB.DataType = os.Getenv("DATABASE_TYPE")
}

func getDSN() string {
	dbType := his.DB.DataType

	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" { // So, if running a container or have a database password; dev this dns...
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s", // otherwise if users have local postgres on machine with no password ; dev this dsn...
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		}
		return dsn
	} //if not postgres, build a differnt DSN...
	return "mysql://" + his.BuildDSN()

}

func showHelp() {
	color.Yellow(`Available commands:

	help			- show the help commands
	version			- print application version
	migrate         - runs all up migrations that have not been run previously
	migrate down    - reverses the most recent migration
	
	`)
}
