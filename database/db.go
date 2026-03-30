package database

import (
	"fallout-vault/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL n'est pas définie")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture de la connexion GORM: %v", err)
	}

	if DB.Migrator().HasTable(&models.Game{}) {
		if err := DB.Exec("UPDATE games SET tags = '' WHERE tags IS NULL").Error; err != nil {
			return fmt.Errorf("erreur lors du nettoyage des tags: %v", err)
		}
	}

	err = DB.AutoMigrate(&models.Game{})
	if err != nil {
		return fmt.Errorf("erreur lors de la migration: %v", err)
	}

	log.Println("✓ Connexion GORM/PostgreSQL établie et migration effectuée")
	return nil
}

func CreateTables() error {
	log.Println("✓ Migration automatique avec GORM, pas besoin de requête SQL")
	return nil
}

func SeedData() error {
	var count int64
	err := DB.Model(&models.Game{}).Count(&count).Error
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification des données: %v", err)
	}

	if count > 0 {
		log.Println("✓ Données déjà présentes, skip seed")
		return nil
	}

	games := []models.Game{
		{
			Title:       "Fallout",
			GameType:    "RPG",
			Description: "Le premier jeu de la franchise Fallout. Un RPG post-apocalyptique isométrique où vous devez sauver votre abri en trouvant une puce d'eau. Considéré comme un classique du genre.",
			ReleaseYear: 1997,
			Developer:   "Interplay Productions",
			Platforms:   "PC, Mac",
			Rating:      9.0,
			ImageURL:    "images/fallout1.jpg",
			Tags:        "",
		},
		{
			Title:       "Fallout 2",
			GameType:    "RPG",
			Description: "La suite directe de Fallout, avec un monde encore plus vaste et des choix plus complexes. Vous incarnez un descendant du Vault Dweller original.",
			ReleaseYear: 1998,
			Developer:   "Black Isle Studios",
			Platforms:   "PC, Mac",
			Rating:      9.2,
			ImageURL:    "images/fallout2.jpg",
			Tags:        "",
		},
		{
			Title:       "Fallout Tactics",
			GameType:    "Tactical RPG",
			Description: "Un jeu tactique spin-off se déroulant dans le Midwest américain. Vous commandez une escouade de la Confrérie de l'Acier dans des missions stratégiques en temps réel ou au tour par tour.",
			ReleaseYear: 2001,
			Developer:   "Micro Forté / 14 Degrees East",
			Platforms:   "PC",
			Rating:      7.5,
			ImageURL:    "images/fallout_tactics.jpg",
			Tags:        "",
		},
		{
			Title:       "Fallout 3",
			GameType:    "Action RPG",
			Description: "Le renouveau de la franchise par Bethesda. Un RPG en vue première personne dans les ruines de Washington D.C. avec le système V.A.T.S.",
			ReleaseYear: 2008,
			Developer:   "Bethesda Game Studios",
			Platforms:   "PC, Xbox 360, PS3",
			Rating:      9.1,
			ImageURL:    "images/fallout3.jpg",
			Tags:        "",
		},
		{
			Title:       "Fallout: New Vegas",
			GameType:    "Action RPG",
			Description: "Développé par Obsidian Entertainment, ce jeu se déroule dans le désert de Mojave. Considéré par beaucoup comme le meilleur Fallout moderne grâce à son écriture exceptionnelle.",
			ReleaseYear: 2010,
			Developer:   "Obsidian Entertainment",
			Platforms:   "PC, Xbox 360, PS3",
			Rating:      9.5,
			ImageURL:    "images/fallout_new_vegas.jpg",
			Tags:        "",
		},
		{
			Title:       "Fallout 4",
			GameType:    "Action RPG",
			Description: "Le dernier RPG solo majeur de la franchise, se déroulant à Boston. Introduit un système de construction de colonies et une personnalisation d'armes avancée.",
			ReleaseYear: 2015,
			Developer:   "Bethesda Game Studios",
			Platforms:   "PC, Xbox One, PS4",
			Rating:      8.5,
			ImageURL:    "images/fallout4.jpg",
			Tags:        "",
		},
		{
			Title:       "Fallout 76",
			GameType:    "Online RPG",
			Description: "Le premier Fallout multijoueur en ligne. Se déroule en Virginie-Occidentale, 25 ans après la guerre nucléaire. Controversial au lancement mais amélioré avec le temps.",
			ReleaseYear: 2018,
			Developer:   "Bethesda Game Studios",
			Platforms:   "PC, Xbox One, PS4",
			Rating:      6.5,
			ImageURL:    "images/fallout76.jpg",
			Tags:        "",
		},
		{
			Title:       "Fallout Shelter",
			GameType:    "Mobile/Strategy",
			Description: "Un jeu de gestion d'abri pour mobiles et PC. Construisez et gérez votre propre Vault-Tec, assignez des dwellers à différentes tâches.",
			ReleaseYear: 2015,
			Developer:   "Bethesda Game Studios",
			Platforms:   "iOS, Android, PC, Xbox, PS4, Switch",
			Rating:      7.8,
			ImageURL:    "images/fallout_shelter.png",
			Tags:        "",
		},
	}

	err = DB.Create(&games).Error
	if err != nil {
		return fmt.Errorf("erreur lors de l'insertion des jeux: %v", err)
	}

	log.Printf("✓ %d jeux Fallout insérés avec succès", len(games))
	return nil
}

func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err == nil {
			sqlDB.Close()
			log.Println("✓ Connexion à la base de données fermée")
		}
	}
}
