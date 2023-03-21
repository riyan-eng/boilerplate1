package config

import (
	"fmt"
	"log"
	"time"

	"boilerplate/migration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=riyan dbname=boilerplate1 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal("can't connect to database")
	}

	sqlDB, _ := Database.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("can't ping to database")
	}

	fmt.Println("connection opened to database")
	Database.AutoMigrate(
		&migration.Companies{}, &migration.UserTypes{}, &migration.Users{}, &migration.UserDatas{}, &migration.Coas{}, &migration.Transactions{}, &migration.Journals{},
	)
}
