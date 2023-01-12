package utils

import (
	"fmt"
	config "rogerdev_golf/configs"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)
	fmt.Println(connectionString, "hallo")
	DB, err := gorm.Open(mysql.Open(connectionString) /* &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true} */)

	if err != nil {
		log.Info("error in connect database : ", err)
		panic(err)
	}

	// AutoMigrate(DB)
	return DB
}
