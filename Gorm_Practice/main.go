package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

func main() {
	// Configure the PostgreSQL Connection
	dsn := "host=localhost user=postgres password=? dbname=GoTest port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// Open the Database Connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established")

	// Run Migrations
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration completed")

	// Crud Operations
	// create
	newUser := User{Name: "John Doe", Email: "john@example.com", Password: "password123"}
	db.Create(&newUser)

	// read
	var user User
	db.First(&user, 1)                               // Find user with primary key 1
	db.First(&user, "email = ?", "john@example.com") // Find user with email
	println(user.Name, user.Email, user.Password)

	// update
	db.Model(&user).Update("Name", "Johnny")

}
