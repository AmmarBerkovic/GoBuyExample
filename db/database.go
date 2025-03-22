package db

import (
	"fmt"
	"log"

	"github.com/AmmarBerkovic/GoBuyExample/config" // Import the config package
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global database instance
var DB *gorm.DB

// ConnectDatabase initializes the MySQL connection
func ConnectDatabase() {
	// Load configuration
	cfg, err := config.LoadConfig() // Get config from config.yaml
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}
	// Get database credentials from config
	dbUser := cfg.Database.User
	dbPassword := cfg.Database.Password
	dbHost := cfg.Database.Host
	dbPort := fmt.Sprintf("%d", cfg.Database.Port)
	dbName := cfg.Database.Name

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// Open a connection to the database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Assign the database instance to the global variable
	DB = database

	// Run auto-migration for the User model
	// err = DB.AutoMigrate(&models.User{})
	// if err != nil {
	// 	log.Fatal("Failed to run migrations:", err)
	// }

	log.Println("Database connection established and migrations completed.")
}
