package driversql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Instance *gorm.DB
}

type DBMysqlOption struct {
	Host     string `deepcopier:"field:Host"`
	Port     string `deepcopier:"field:Port"`
	DB       string `deepcopier:"field:DB"`
	Username string `deepcopier:"field:Username"`
	Password string `deepcopier:"field:Password"`
}

func NewMysqlDatabase(option DBMysqlOption) (database *Database, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		option.Username,
		option.Password,
		option.Host,
		option.Port,
		option.DB,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return
	}

	database = &Database{
		Instance: db,
	}

	return
}
