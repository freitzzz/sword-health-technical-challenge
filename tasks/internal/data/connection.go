package data

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	mySQLUsernameEnvKey = "mysql_user"
	mySQLPasswordEnvKey = "mysql_pass"
	mySQLAddressEnvKey  = "mysql_address"
	mySQLDatabaseEnvKey = "mysql_database"
)

var (
	mySQLUsername = os.Getenv(mySQLUsernameEnvKey)
	mySQLPassword = os.Getenv(mySQLPasswordEnvKey)
	mySQLAddress  = os.Getenv(mySQLAddressEnvKey)
	mySQLDatabase = os.Getenv(mySQLDatabaseEnvKey)
)

func OpenDbConnection() (*gorm.DB, error) {

	// Charset utf8mb4 for fully support on utf-8, including emojis
	// More info: https://gorm.io/docs/connecting_to_the_database.html
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name%20for%20details

	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mySQLUsername, mySQLPassword, mySQLAddress, mySQLDatabase),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
}
