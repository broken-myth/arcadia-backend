package database

import (
	"fmt"
	"strconv"

	"github.com/delta/arcadia-backend/config"
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Function to Connect to the Database
func ConnectMySQLdb() {
	config := config.GetConfig()

	dbPort := strconv.FormatUint(uint64(config.Db.Port), 10)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Db.Username,
		config.Db.Password,
		config.Db.Host,
		dbPort,
		config.Db.Name,
	)

	// Logger for Database, Logs will be saved in config.Db.LogFile
	initDatabaseLogger(config.Db.LogFile)

	var err error

	dblogger := newDBLogger()

	// MySQL connection is established
	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: dblogger,
	})

	if err != nil {
		fmt.Println(err)
		panic(color.RedString("Failed to connect to MySQL database!"))
	} else {
		fmt.Println(color.GreenString("MySQL Database connected!"))
	}
}
