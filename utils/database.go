package utils

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// 使用 .env 檔案中的 SUPABASE_DB_URL 環境變數來建立連接
	dsn := os.Getenv("SUPABASE_DB_URL")
	log.Println("Connecting to database with URL:", dsn)
	if dsn == "" {
		log.Fatal("❌ SUPABASE_DB_URL is not set in the environment variables")
		return
	}

	// 嘗試使用 GORM 連接到 PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to Supabase DB:", err)
		return
	}

	log.Println("✅ Successfully connected to Supabase DB!")
	DB = db
}
