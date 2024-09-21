package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=postgres password=password dbname=localmarketplace port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    } else {
        log.Println("Database connected!")
    }
}
