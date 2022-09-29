package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func OpenConnection() *gorm.DB {
	configDb := "host=localhost user=postgres password=postgres dbname=upvote_project port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(configDb), &gorm.Config{})

	if err != nil {
		fmt.Println("Couldn't connect to database")
		log.Fatal("Error: ", err)
	}

	return db
}

func CloseConnection(connection *gorm.DB) {
	db, err := connection.DB()
	if err != nil {
		fmt.Println("Couldn't close connection to database")
		log.Fatalln(err)
	}

	if err = db.Close(); err != nil {
		fmt.Println("Couldn't close connection to database")
		log.Fatalln(err)
	}
}

//func Migrations(connection *gorm.DB) {
//	connection.AutoMigrate(models.Currency{})
//	connection.AutoMigrate(models.Vote{})
//}
