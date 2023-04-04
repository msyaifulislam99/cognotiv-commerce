package configuration

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"syaiful.com/simple-commerce/entity"
	"syaiful.com/simple-commerce/exception"
)

func NewDatabase(config Config) *gorm.DB {
	dsn := "host=postgres user=postgres password=postgrespassword dbname=commerce port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	exception.PanicLogging(err)

	//autoMigrate
	err = db.AutoMigrate(&entity.User{})
	err = db.AutoMigrate(&entity.UserRole{})
	err = db.AutoMigrate(&entity.Product{})
	err = db.AutoMigrate(&entity.Order{})
	err = db.AutoMigrate(&entity.OrderDetail{})
	exception.PanicLogging(err)
	return db
}
