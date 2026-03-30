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

	if err := database.CreateTables(); err != nil {
		log.Fatalf("Erreur lors de la création des tables: %v", err)
	}

	if err := database.SeedData(); err != nil {
		log.Fatalf("Erreur lors de l'insertion des données: %v", err)
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	router.Static("/static", "./static")

	router.GET("/", handlers.HomeHandler)
	router.GET("/games", handlers.GamesListHandler)
	router.GET("/games/:id", handlers.GameDetailHandler)
	router.GET("/admin/games", handlers.AdminGamesHandler)
	router.GET("/admin/games/new", handlers.AdminGameNewHandler)
	router.POST("/admin/games", handlers.AdminGameCreateHandler)
	router.GET("/admin/games/:id/edit", handlers.AdminGameEditHandler)
	router.POST("/admin/games/:id/update", handlers.AdminGameUpdateHandler)
	router.POST("/admin/games/:id/delete", handlers.AdminGameDeleteHandler)
	router.POST("/admin/games/clear-tags", handlers.AdminGameClearTagsHandler)
	router.GET("/credits", handlers.CreditsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Serveur Fallout Vault demarre sur le port %s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}
