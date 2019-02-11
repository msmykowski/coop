package db

import (
  "api/db/models"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Connection = gorm.DB

func Open() *Connection {
  db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=coop_dev sslmode=disable")

  if err != nil {
    panic(err)
  }

  db.AutoMigrate(&models.JobDefinition{})

  return db
}
