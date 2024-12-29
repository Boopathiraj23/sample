package repository

import (
	"fmt"
	"log"
	"os"
	"todo/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Dbdata *gorm.DB

func Dbconnection() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error occur when loading a .env file : %v", err)
	}
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("HOST"), os.Getenv("PORT"))
	Dbdata, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Fatalf("failed to open database : %v", err)
	}
	if err = Dbdata.AutoMigrate(&model.Todo{}); err != nil {
		log.Fatalf("failed to table creation : %v", err)
	}
	fmt.Println("Table created successfully")

}
