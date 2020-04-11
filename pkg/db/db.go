package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func ConnectPGDB() *gorm.DB {
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		viper.GetString("database.postgres.host"),
		viper.GetString("database.postgres.port"),
		viper.GetString("database.postgres.user"),
		viper.GetString("database.postgres.dbname"),
		viper.GetString("database.postgres.password"),
	)
	db, err := gorm.Open("postgres", connection)
	if err != nil{
		panic(err.Error())
	}
	//log.Print("Database connected successfully")
	return db
}
