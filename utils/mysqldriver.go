package utils

import (
	"database/sql"
	"fmt"
	"rogerdev_golf/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

func InitDB(config *configs.AppConfig) *sql.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Info("error in connect database : ", err)
		panic(err)
	}
	// defer db.Close()

	// fmt.Println(connectionString, "hallo")
	// DB, err := gorm.Open(mysql.Open(connectionString) /* &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true} */)

	// AutoMigrate(DB)
	return db
}
