package handlers

import (
	"fallout-vault/database"
	"fallout-vault/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	var stats struct {
		TotalGames   int
		EarliestYear int
		LatestYear   int
		AvgRating    float64
	}

	var games []models.Game
	if err := database.DB.Find(&games).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "home.html", gin.H{
			"error": "Erreur lors du chargement des statistiques",
		})
		return
	}

	stats.TotalGames = len(games)
	stats.EarliestYear = 9999
	stats.LatestYear = 0
	var totalRating float64
	for _, g := range games {
		if g.ReleaseYear < stats.EarliestYear {
			stats.EarliestYear = g.ReleaseYear
		}
		if g.ReleaseYear > stats.LatestYear {
			stats.LatestYear = g.ReleaseYear
		}
		totalRating += g.Rating
	}
	if stats.TotalGames > 0 {
		stats.AvgRating = totalRating / float64(stats.TotalGames)
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"stats": stats,
	})
}

func GamesListHandler(c *gin.Context) {
	gameTypeFilter := c.Query("type")
	yearFilter := c.Query("year")

	var games []models.Game
	query := database.DB.Order("release_year ASC")

	if gameTypeFilter != "" {
		query = query.Where("game_type = ?", gameTypeFilter)
	}
	if yearFilter != "" {
		query = query.Where("release_year = ?", yearFilter)
	}

	if err := query.Find(&games).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "games.html", gin.H{
			"error": "Erreur lors du chargement des jeux",
		})
		return
	}

	gameTypes := getGameTypes()
	years := getYears()

	c.HTML(http.StatusOK, "games.html", gin.H{
		"games":      games,
		"gameTypes":  gameTypes,
		"years":      years,
		"typeFilter": gameTypeFilter,
		"yearFilter": yearFilter,
	})
}

func GameDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "game_detail.html", gin.H{
			"error": "ID invalide",
		})
		return
	}

	var game models.Game
	if err := database.DB.First(&game, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "game_detail.html", gin.H{
			"error": "Jeu non trouve",
		})
		return
	}

	ingameImages := getIngameImages(game.ID)

	c.HTML(http.StatusOK, "game_detail.html", gin.H{
		"game":         game,
		"ingameImages": ingameImages,
	})
}

func getIngameImages(gameID int) []string {
	gameFiles := map[int]struct {
		base string
		ext  string
	}{
		1: {"fallout1", ".jpg"},
		2: {"fallout2", ".jpg"},
		3: {"fallout3", ".jpg"},
		4: {"fallout_new_vegas", ".jpg"},
		5: {"fallout4", ".jpg"},
		6: {"fallout76", ".jpg"},
		7: {"fallout_shelter", ".jpg"},
	}

	data, exists := gameFiles[gameID]
	if !exists {
		return []string{}
	}

	return []string{
		"/static/images_ingames/" + data.base + "_ingame1" + data.ext,
		"/static/images_ingames/" + data.base + "_ingame2" + data.ext,
		"/static/images_ingames/" + data.base + "_ingame3" + data.ext,
	}
}

func getGameTypes() []string {
	var types []string
	var games []models.Game
	database.DB.Distinct("game_type").Order("game_type").Find(&games)
	typeSet := make(map[string]bool)
	for _, g := range games {
		if !typeSet[g.GameType] {
			types = append(types, g.GameType)
			typeSet[g.GameType] = true
		}
	}
	return types
}

func getYears() []int {
	var years []int
	var games []models.Game
	database.DB.Distinct("release_year").Order("release_year DESC").Find(&games)
	yearSet := make(map[int]bool)
	for _, g := range games {
		if !yearSet[g.ReleaseYear] {
			years = append(years, g.ReleaseYear)
			yearSet[g.ReleaseYear] = true
		}
	}
	return years
}
