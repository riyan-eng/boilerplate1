package migration

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Create() {
	dsn := "host=localhost user=postgres password=riyan dbname=boilerplate1 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal("migration: can't connect to database")
	}

	sqlDB, _ := database.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("migration: can't ping to database")
	}

	fmt.Println("migration: connection opened to database")
	fmt.Println("migration: start")
	database.AutoMigrate(
		&Companies{}, &UserTypes{}, &Users{}, &UserDatas{}, &Coas{}, &Transactions{}, &Journals{},
	)
	fmt.Println("migration: done")
	sqlDB.Close()
}
