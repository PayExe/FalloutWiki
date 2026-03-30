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
	// Chargement du fichier .env
	if err := godotenv.Load(); err != nil {
		log.Println("Aucun fichier .env trouvé, on continue...")
	}

	// Initialisation de la base de données
	if err := database.InitDB(); err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la DB: %v", err)
	}
	defer database.CloseDB()

	// Création des tables
	if err := database.CreateTables(); err != nil {
		log.Fatalf("Erreur lors de la création des tables: %v", err)
	}

	// Insertion des données de test
	if err := database.SeedData(); err != nil {
		log.Fatalf("Erreur lors de l'insertion des données: %v", err)
	}

	// Configuration de Gin
	router := gin.Default()

	// Chargement des templates HTML
	router.LoadHTMLGlob("templates/*.html")

	// Servir les fichiers statiques
	router.Static("/static", "./static")

	// Routes
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

	// Récupération du port depuis les variables d'environnement (Scalingo)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port par défaut pour le développement local
	}

	log.Printf("🚀 Serveur Fallout Vault démarré sur le port %s", port)
	log.Printf("📍 Accédez à http://localhost:%s", port)

	// Démarrage du serveur
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}
