package repositories

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"trungpham/gowebbasic/config"
)

func ConnectMysqlInit() *gorm.DB {
	var mysqlVar = config.GetMysqlEnvConfig()

	dsn := mysqlVar.User + ":" + mysqlVar.Password + "@tcp(" + mysqlVar.Host + ":" + mysqlVar.Port + ")/" + mysqlVar.DB + "?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Open DB: ", err)
	}
	return db
}
