package app

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func Dbconncentio() *gorm.DB {
	dbHost := os.Getenv("DBHOST")
	dbUsername := os.Getenv("DBUSERNAME")
	dbPassword := os.Getenv("DBPASSWORD")
	dbDatabase := os.Getenv("DBDATABASE")
	dbPort 	:= os.Getenv("DBPORT")

	dsn := fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
    dbHost, dbPort, dbUsername, dbPassword, dbDatabase,
)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err.Error())
	}
	return  db
}