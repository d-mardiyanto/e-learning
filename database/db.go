package database

import (
	"e-learning/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	// Build connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	// Co

	// Connect to the database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// // Migrate models
	DB.AutoMigrate(
		&models.Instructors{},
		&models.Students{},
		&models.Students_Academic{},
		&models.Course{},
		&models.CourseFiles{},
		&models.Classes{},
		&models.StudyProgram{},
		&models.Quiz{},
		&models.Questions{},
		&models.Answer{},
		&models.Answer_Detail{},
		&models.Transaction{},
		&models.TransactionDetail{},
	)
	log.Println("Database connection established")
}
