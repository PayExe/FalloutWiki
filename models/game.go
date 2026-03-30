package models

import "time"

// Game représente un jeu Fallout dans la base de données
type Game struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	GameType    string    `gorm:"size:100;not null" json:"game_type"`
	Description string    `gorm:"type:text" json:"description"`
	ReleaseYear int       `gorm:"not null" json:"release_year"`
	Developer   string    `gorm:"size:255;not null" json:"developer"`
	Platforms   string    `gorm:"size:255" json:"platforms"`
	Rating      float64   `gorm:"type:decimal(3,1)" json:"rating"`
	ImageURL    string    `gorm:"size:500" json:"image_url"`
	Tags        string    `gorm:"type:text;not null;default:''" json:"tags"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
