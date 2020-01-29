package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Driver is a syntax sugar for database driver
type Driver = *gorm.DB

// New starts a new database connection with received credentials
func New(host, user, dbname, password string) (Driver, error) {
	return gorm.Open(
		"postgres",
		"host="+host+" user="+user+
			" dbname="+dbname+" sslmode=disable password="+
			password)
}
