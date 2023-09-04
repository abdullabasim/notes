package database

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	once sync.Once
	db   *sql.DB
)

func InitDB() *sql.DB {
	once.Do(func() {

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("config")

		// Read the configuration file
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Error reading config file:", err)
		}

		// Get database configuration from Viper
		dbUser := viper.GetString("database.user")
		dbPassword := viper.GetString("database.password")
		dbName := viper.GetString("database.dbname")
		dbHost := viper.GetString("database.host")
		dbPort := viper.GetString("database.port")
		dbSSLMode := viper.GetString("database.sslmode")

		connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
			dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode)

		// Create a new connection
		conn, err := sqlx.Connect("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}

		// Assign the database instance
		db = conn.DB

		// Create the notes table if it doesn't exist
		createTableQuery := `
            CREATE TABLE IF NOT EXISTS notes (
                id SERIAL PRIMARY KEY,
                title TEXT,
                text TEXT,
                created_at TIMESTAMP DEFAULT NOW(),
                updated_at TIMESTAMP
            )
        `
		_, err = db.Exec(createTableQuery)
		if err != nil {
			log.Fatal(err)
		}
	})

	return db
}
