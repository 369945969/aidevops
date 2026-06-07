package main

import (
	"log"
	"os"

	"devops/handlers"
	"devops/models"
	"devops/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "db/devops.db"
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get underlying sql.DB: %v", err)
	}
	sqlDB.SetMaxOpenConns(1)

	err = db.AutoMigrate(
		&models.AutomationConfig{},
		&models.ReviewNode{},
		&models.AITeam{},
		&models.AIAgent{},
		&models.AIModel{},
		&models.CodeRepo{},
		&models.CICDService{},
		&models.CloudCredential{},
		&models.SSHKey{},
		&models.EnvVar{},
		&models.ConnectedService{},
		&models.APIKey{},
		&models.WebhookConfig{},
		&models.SyncConfig{},
		&models.NotificationChannel{},
		&models.NotificationEvent{},
		&models.NotificationPreference{},
		&models.ReviewRule{},
		&models.ReviewTemplate{},
		&models.ReviewTimeoutConfig{},
		&models.Role{},
		&models.Permission{},
		&models.RolePermission{},
		&models.SecurityPolicy{},
		&models.PasswordPolicy{},
		&models.AuditLog{},
		&models.TeamMember{},
		&models.PendingInvite{},
		&models.MemberAITeamBinding{},
	)
	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	log.Println("Database migrated successfully")

	seedDB(db)

	h := handlers.NewHandler(db)

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	routes.SetupRoutes(r, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}