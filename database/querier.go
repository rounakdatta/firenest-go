package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetConnection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	databaseCredentials := fmt.Sprintf("%s:%s@/%s", os.Getenv("APP_USER"), os.Getenv("APP_PASSWORD"), os.Getenv("APP_DATABASE"))
	db, err := sql.Open("mysql", databaseCredentials)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAccountIdFromName(db *sql.DB, accountName string) int {
	defer db.Close()

	query := fmt.Sprintf("SELECT id FROM accounts WHERE name='%s' AND deleted_at IS NULL LIMIT 1", accountName)
	queryOutput := db.QueryRow(query)

	var accountId int
	err := queryOutput.Scan(&accountId)

	if err != nil {
		return -1
	}

	return accountId
}
