package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend/internal/dsn"
	"backend/internal/model"
)

func main() {
    _ = godotenv.Load()
    db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Явно мигрировать только нужные таблицы
    err = db.AutoMigrate(&model.Ship{},&model.Request{}, &model.User{}, &model.RequestShip{})
    if err != nil {
        panic("cant migrate db")
    }
}
