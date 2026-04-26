package main

import (
	"announcement-board/handlers"
	"announcement-board/models"
	"announcement-board/repositories"
	"announcement-board/services"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// เชื่อมต่อ Database
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=announcements port=5432 sslmode=disable"
	}

	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("DB not ready, retrying... (%d/10)", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	db.AutoMigrate(&models.Announcement{})
	log.Println("Database connected and migrated!")

	// Wire up layers
	repo := repositories.NewAnnouncementRepository(db)
	service := services.NewAnnouncementService(repo)
	handler := handlers.NewAnnouncementHandler(service)

	// Setup Router
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Routes
	r.GET("/announcements", handler.GetAll)
	r.POST("/announcements", handler.CreateAnnouncementRequest)
	r.PUT("/announcements/:id", handler.UpdateAnnouncementRequest)
	r.DELETE("/announcements/:id", handler.DeleteAnnouncementRequest)

	port := os.Getenv("PORT")
	if port == "" {
    	port = "8080"
	}
	
	r.Run(":" + port)
}

