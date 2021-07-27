package data

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	_MySQLUsernameEnvKey = "mysql.user"
	_MySQLPasswordEnvKey = "mysql.pass"
	_MySQLAddressEnvKey  = "mysql.address"
	_MySQLDatabaseEnvKey = "mysql.database"
)

var (
	_MySQLUsername = os.Getenv(_MySQLUsernameEnvKey)
	_MySQLPassword = os.Getenv(_MySQLPasswordEnvKey)
	_MySQLAddress  = os.Getenv(_MySQLAddressEnvKey)
	_MySQLDatabase = os.Getenv(_MySQLDatabaseEnvKey)
)

func OpenDbConnection() (*gorm.DB, error) {

	// Charset utf8mb4 for fully support on utf-8, including emojis
	// More info: https://gorm.io/docs/connecting_to_the_database.html
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name%20for%20details

	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", _MySQLUsername, _MySQLPassword, _MySQLAddress, _MySQLDatabase),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
}
