package database

import (
	"database/sql"
	"fmt"
	"github.com/rounakdatta/firenest/utils"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetConnection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	databaseCredentials := fmt.Sprintf("%s:%s@%s/%s", os.Getenv("APP_USER"), os.Getenv("APP_PASSWORD"), os.Getenv("APP_DATABASE_HOST_FORMATTED"), os.Getenv("APP_DATABASE"))
	db, err := sql.Open("mysql", databaseCredentials)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAccountIdFromName(db *sql.DB, accountName string, direction utils.Direction) int {
	defer db.Close()

	accountType := "Expense account"
	if direction == utils.CREDIT {
		accountType = "Revenue account"
	} else if direction == utils.OWN {
		accountType = "Asset account"
	}

	query := fmt.Sprintf("SELECT id FROM accounts WHERE name='%s' AND deleted_at IS NULL AND account_type_id IN ((SELECT id FROM account_types WHERE type='%s')) LIMIT 1", accountName, accountType)
	queryOutput := db.QueryRow(query)

	var accountId int
	err := queryOutput.Scan(&accountId)

	if err != nil {
		return -1
	}

	return accountId
}
