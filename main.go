package main

import (
	"fallout-vault/database"
	"fallout-vault/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aucun fichier .env trouvé, on continue...")
	}

	if err := database.InitDB(); err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la DB: %v", err)
	}
	defer database.CloseDB()

	if err := database.SeedData(); err != nil {
		log.Fatalf("Erreur lors de l'insertion des données: %v", err)
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	router.Static("/static", "./static")

	router.GET("/", handlers.HomeHandler)
	router.GET("/games", handlers.GamesListHandler)
	router.GET("/games/:id", handlers.GameDetailHandler)
	router.GET("/admin", handlers.AdminLoginHandler)
	router.POST("/admin", handlers.AdminLoginSubmitHandler)
	router.GET("/admin/logout", handlers.AdminLogoutHandler)
	admin := router.Group("/admin")
	admin.Use(handlers.AdminAuthMiddleware())
	admin.GET("/games", handlers.AdminGamesHandler)
	admin.GET("/games/new", handlers.AdminGameNewHandler)
	admin.POST("/games", handlers.AdminGameCreateHandler)
	admin.GET("/games/:id/edit", handlers.AdminGameEditHandler)
	admin.POST("/games/:id/update", handlers.AdminGameUpdateHandler)
	admin.POST("/games/:id/delete", handlers.AdminGameDeleteHandler)
	admin.POST("/games/clear-tags", handlers.AdminGameClearTagsHandler)
	router.GET("/credits", handlers.CreditsHandler)
	router.GET("/health", handlers.HealthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Serveur Fallout Vault demarre sur le port %s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}
