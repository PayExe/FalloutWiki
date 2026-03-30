package handlers

import (
	"fallout-vault/database"
	"fallout-vault/models"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const bethesdaURL = "https://bethesda.net/"

type PageStats struct {
	TotalGames    int
	DistinctTypes int
	EarliestYear  int
	LatestYear    int
	BethesdaGames int
	AverageRating float64
}

type CatalogGame struct {
	models.Game
	Tags          []string
	TagLine       string
	ScreenshotURL string
}

type HomePageData struct {
	ActivePage    string
	ContentPage   string
	PageTitle     string
	Error         string
	Stats         PageStats
	FeaturedGames []CatalogGame
	BethesdaURL   string
}

type CatalogPageData struct {
	ActivePage   string
	ContentPage  string
	PageTitle    string
	Error        string
	Stats        PageStats
	Games        []CatalogGame
	AllTags      []string
	AllTypes     []string
	AllYears     []int
	SelectedTag  string
	SelectedType string
	SelectedYear string
	SearchQuery  string
	ResultCount  int
	TotalCount   int
	BethesdaURL  string
}

type DetailFact struct {
	Label string
	Value string
}

type DetailPageData struct {
	ActivePage   string
	ContentPage  string
	PageTitle    string
	Error        string
	Game         CatalogGame
	Tags         []string
	Facts        []DetailFact
	Screenshots  []string
	RelatedGames []CatalogGame
	BethesdaURL  string
}

type CreditsPageData struct {
	ActivePage  string
	ContentPage string
	PageTitle   string
	TechStack   []string
	Reasons     []string
	Notes       []string
	BethesdaURL string
}

type AdminPageData struct {
	ActivePage  string
	ContentPage string
	PageTitle   string
	Error       string
	Games       []CatalogGame
	BethesdaURL string
}

type GameFormPageData struct {
	ActivePage  string
	ContentPage string
	PageTitle   string
	Error       string
	Game        models.Game
	ActionURL   string
	IsEdit      bool
	BethesdaURL string
}

func HomeHandler(c *gin.Context) {
	games, err := loadGames()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "home.html", HomePageData{
			ActivePage:  "home",
			ContentPage: "home_content",
			PageTitle:   "Fallout Vault - Accueil",
			Error:       "Impossible de charger le site",
		})
		return
	}

	stats := computeStats(games)
	featured := bestGames(games, 3)

	c.HTML(http.StatusOK, "home.html", HomePageData{
		ActivePage:    "home",
		ContentPage:   "home_content",
		PageTitle:     "Fallout Vault - Accueil",
		Stats:         stats,
		FeaturedGames: featured,
		BethesdaURL:   bethesdaURL,
	})
}

func GamesListHandler(c *gin.Context) {
	games, err := loadGames()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "games.html", CatalogPageData{
			ActivePage:  "catalog",
			ContentPage: "games_content",
			PageTitle:   "Fallout Vault - Catalogue",
			Error:       "Impossible de charger le catalogue",
		})
		return
	}

	selectedTag := normalizeFilter(c.Query("tag"))
	selectedType := normalizeFilter(c.Query("type"))
	selectedYear := normalizeFilter(c.Query("year"))
	searchQuery := strings.TrimSpace(c.Query("q"))

	filtered := filterGames(games, selectedTag, selectedType, selectedYear, searchQuery)
	stats := computeStats(games)

	c.HTML(http.StatusOK, "games.html", CatalogPageData{
		ActivePage:   "catalog",
		ContentPage:  "games_content",
		PageTitle:    "Fallout Vault - Catalogue",
		Stats:        stats,
		Games:        filtered,
		AllTags:      collectTags(games),
		AllTypes:     collectTypes(games),
		AllYears:     collectYears(games),
		SelectedTag:  selectedTag,
		SelectedType: selectedType,
		SelectedYear: selectedYear,
		SearchQuery:  searchQuery,
		ResultCount:  len(filtered),
		TotalCount:   len(games),
		BethesdaURL:  bethesdaURL,
	})
}

func GameDetailHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "game_detail.html", DetailPageData{
			ActivePage:  "catalog",
			ContentPage: "game_detail_content",
			PageTitle:   "Fallout Vault - Dossier",
			Error:       "ID invalide",
		})
		return
	}

	games, err := loadGames()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "game_detail.html", DetailPageData{
			ActivePage:  "catalog",
			ContentPage: "game_detail_content",
			PageTitle:   "Fallout Vault - Dossier",
			Error:       "Impossible de charger la fiche",
		})
		return
	}

	current, ok := findGameByID(games, id)
	if !ok {
		c.HTML(http.StatusNotFound, "game_detail.html", DetailPageData{
			ActivePage:  "catalog",
			ContentPage: "game_detail_content",
			PageTitle:   "Fallout Vault - Dossier",
			Error:       "Jeu non trouvé",
		})
		return
	}

	facts := []DetailFact{
		{Label: "ID DB", Value: strconv.Itoa(current.ID)},
		{Label: "Type", Value: current.GameType},
		{Label: "Année", Value: strconv.Itoa(current.ReleaseYear)},
		{Label: "Développeur", Value: current.Developer},
		{Label: "Plateformes", Value: current.Platforms},
		{Label: "Note", Value: strconv.FormatFloat(current.Rating, 'f', 1, 64)},
		{Label: "Tags", Value: strings.Join(current.Tags, ", ")},
	}

	c.HTML(http.StatusOK, "game_detail.html", DetailPageData{
		ActivePage:   "catalog",
		ContentPage:  "game_detail_content",
		PageTitle:    current.Title + " - Fallout Vault",
		Game:         current,
		Tags:         current.Tags,
		Facts:        facts,
		Screenshots:  getIngameImages(current.ID),
		RelatedGames: relatedGames(games, current),
		BethesdaURL:  bethesdaURL,
	})
}

func CreditsHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "credits.html", CreditsPageData{
		ActivePage:  "credits",
		ContentPage: "credits_content",
		PageTitle:   "Fallout Vault - Crédits",
		TechStack:   []string{"Go", "Gin", "GORM", "PostgreSQL", "HTML", "CSS", "Vanilla JS"},
		Reasons: []string{
			"Un site vitrine pensé comme un journal de survie dans les ruines.",
			"Un catalogue clair pour parcourir toute la licence Fallout en un coup d'oeil.",
			"Une vraie page détail façon dossier de terrain, avec visuels et infos DB.",
		},
		Notes: []string{
			"Fallout m'accompagne depuis l'enfance et ce site est une lettre d'amour à la licence.",
			"Projet non officiel, réalisé pour le plaisir et la culture jeu vidéo.",
		},
		BethesdaURL: bethesdaURL,
	})
}

func AdminGamesHandler(c *gin.Context) {
	games, err := loadGames()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_games.html", AdminPageData{
			ActivePage:  "admin",
			ContentPage: "admin_games_content",
			PageTitle:   "Fallout Vault - Admin",
			Error:       "Impossible de charger les jeux",
			BethesdaURL: bethesdaURL,
		})
		return
	}

	c.HTML(http.StatusOK, "admin_games.html", AdminPageData{
		ActivePage:  "admin",
		ContentPage: "admin_games_content",
		PageTitle:   "Fallout Vault - Admin",
		Games:       games,
		BethesdaURL: bethesdaURL,
	})
}

func AdminGameNewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_game_form.html", GameFormPageData{
		ActivePage:  "admin",
		ContentPage: "admin_game_form_content",
		PageTitle:   "Fallout Vault - Nouveau jeu",
		ActionURL:   "/admin/games",
		BethesdaURL: bethesdaURL,
	})
}

func AdminGameEditHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_game_form.html", GameFormPageData{
			ActivePage:  "admin",
			ContentPage: "admin_game_form_content",
			PageTitle:   "Fallout Vault - Editer",
			Error:       "ID invalide",
			BethesdaURL: bethesdaURL,
		})
		return
	}

	var game models.Game
	if err := database.DB.First(&game, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "admin_game_form.html", GameFormPageData{
			ActivePage:  "admin",
			ContentPage: "admin_game_form_content",
			PageTitle:   "Fallout Vault - Editer",
			Error:       "Jeu introuvable",
			BethesdaURL: bethesdaURL,
		})
		return
	}

	c.HTML(http.StatusOK, "admin_game_form.html", GameFormPageData{
		ActivePage:  "admin",
		ContentPage: "admin_game_form_content",
		PageTitle:   "Fallout Vault - Editer " + game.Title,
		Game:        game,
		ActionURL:   "/admin/games/" + strconv.Itoa(game.ID) + "/update",
		IsEdit:      true,
		BethesdaURL: bethesdaURL,
	})
}

func AdminGameCreateHandler(c *gin.Context) {
	game, err := gameFromForm(c)
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_game_form.html", GameFormPageData{
			ActivePage:  "admin",
			ContentPage: "admin_game_form_content",
			PageTitle:   "Fallout Vault - Nouveau jeu",
			Error:       err.Error(),
			Game:        game,
			ActionURL:   "/admin/games",
			BethesdaURL: bethesdaURL,
		})
		return
	}

	if err := database.DB.Create(&game).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "admin_game_form.html", GameFormPageData{
			ActivePage:  "admin",
			ContentPage: "admin_game_form_content",
			PageTitle:   "Fallout Vault - Nouveau jeu",
			Error:       "Impossible de créer le jeu",
			Game:        game,
			ActionURL:   "/admin/games",
			BethesdaURL: bethesdaURL,
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin/games")
}

func AdminGameUpdateHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "admin_game_form.html", GameFormPageData{
			ActivePage:  "admin",
			ContentPage: "admin_game_form_content",
			PageTitle:   "Fallout Vault - Editer",
			Error:       "ID invalide",
			BethesdaURL: bethesdaURL,
		})
		return
	}

	game, err := gameFromForm(c)
	if err != nil {
		game.ID = id
		c.HTML(http.StatusBadRequest, "admin_game_form.html", GameFormPageData{
			ActivePage:  "admin",
			ContentPage: "admin_game_form_content",
			PageTitle:   "Fallout Vault - Editer",
			Error:       err.Error(),
			Game:        game,
			ActionURL:   "/admin/games/" + strconv.Itoa(id) + "/update",
			IsEdit:      true,
			BethesdaURL: bethesdaURL,
		})
		return
	}

	updates := map[string]any{
		"title":        game.Title,
		"game_type":    game.GameType,
		"description":  game.Description,
		"release_year": game.ReleaseYear,
		"developer":    game.Developer,
		"platforms":    game.Platforms,
		"rating":       game.Rating,
		"image_url":    game.ImageURL,
		"tags":         game.Tags,
	}
	if err := database.DB.Model(&models.Game{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		game.ID = id
		c.HTML(http.StatusInternalServerError, "admin_game_form.html", GameFormPageData{
			ActivePage:  "admin",
			ContentPage: "admin_game_form_content",
			PageTitle:   "Fallout Vault - Editer",
			Error:       "Impossible de mettre à jour le jeu",
			Game:        game,
			ActionURL:   "/admin/games/" + strconv.Itoa(id) + "/update",
			IsEdit:      true,
			BethesdaURL: bethesdaURL,
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin/games")
}

func AdminGameDeleteHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide")
		return
	}

	if err := database.DB.Delete(&models.Game{}, id).Error; err != nil {
		c.String(http.StatusInternalServerError, "Impossible de supprimer le jeu")
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin/games")
}

func AdminGameClearTagsHandler(c *gin.Context) {
	if err := database.DB.Model(&models.Game{}).Where("1 = 1").Update("tags", "").Error; err != nil {
		c.String(http.StatusInternalServerError, "Impossible de vider les tags")
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin/games")
}

func gameFromForm(c *gin.Context) (models.Game, error) {
	releaseYear, err := strconv.Atoi(strings.TrimSpace(c.PostForm("release_year")))
	if err != nil {
		return models.Game{}, fmt.Errorf("l'année de sortie est invalide")
	}

	rating, err := strconv.ParseFloat(strings.TrimSpace(c.PostForm("rating")), 64)
	if err != nil {
		return models.Game{}, fmt.Errorf("la note est invalide")
	}

	game := models.Game{
		Title:       strings.TrimSpace(c.PostForm("title")),
		GameType:    strings.TrimSpace(c.PostForm("game_type")),
		Description: strings.TrimSpace(c.PostForm("description")),
		ReleaseYear: releaseYear,
		Developer:   strings.TrimSpace(c.PostForm("developer")),
		Platforms:   strings.TrimSpace(c.PostForm("platforms")),
		Rating:      rating,
		ImageURL:    strings.TrimSpace(c.PostForm("image_url")),
		Tags:        normalizeTagsInput(c.PostForm("tags")),
	}

	if game.Title == "" || game.GameType == "" || game.Description == "" || game.Developer == "" || game.Platforms == "" || game.ImageURL == "" {
		return game, fmt.Errorf("merci de remplir tous les champs")
	}

	return game, nil
}

func normalizeTagsInput(raw string) string {
	parts := splitTags(raw)
	if len(parts) == 0 {
		return ""
	}
	return strings.Join(parts, ", ")
}

func loadGames() ([]CatalogGame, error) {
	var rawGames []models.Game
	if err := database.DB.Order("release_year ASC, title ASC").Find(&rawGames).Error; err != nil {
		return nil, err
	}

	games := make([]CatalogGame, 0, len(rawGames))
	for _, game := range rawGames {
		game.ImageURL = resolveAssetURL(game.ImageURL)
		g := CatalogGame{Game: game}
		g.Tags = buildTags(game)
		g.TagLine = strings.Join(g.Tags, " • ")
		games = append(games, g)
	}

	return games, nil
}

func buildTags(game models.Game) []string {
	base := splitTags(game.Tags)

	gameType := normalizeFilter(game.GameType)
	seen := map[string]bool{}
	result := make([]string, 0, len(base))
	for _, tag := range base {
		tag = normalizeFilter(tag)
		if tag == "" || tag == gameType || seen[tag] {
			continue
		}
		seen[tag] = true
		result = append(result, tag)
		if len(result) == 3 {
			break
		}
	}

	return result
}

func splitTags(raw string) []string {
	if strings.TrimSpace(raw) == "" {
		return nil
	}

	parts := strings.FieldsFunc(raw, func(r rune) bool {
		return r == ',' || r == ';' || r == '|'
	})
	cleaned := make([]string, 0, len(parts))
	for _, part := range parts {
		part = normalizeFilter(part)
		if part != "" {
			cleaned = append(cleaned, part)
		}
	}
	return cleaned
}

func derivedTags(game models.Game) []string {
	byTitle := map[string][]string{
		"fallout":            {"classic", "isometric", "vault", "wasteland"},
		"fallout 2":          {"classic", "isometric", "black isle", "wasteland"},
		"fallout tactics":    {"tactical", "squad", "brotherhood", "strategy"},
		"fallout 3":          {"bethesda", "vats", "capital wasteland"},
		"fallout: new vegas": {"obsidian", "mojave", "dialogue", "factions"},
		"fallout 4":          {"bethesda", "settlements", "crafting", "commonwealth"},
		"fallout 76":         {"bethesda", "multiplayer", "online", "appalachia"},
		"fallout shelter":    {"management", "vault", "mobile", "strategy"},
	}

	title := strings.ToLower(game.Title)
	tags := []string{strings.ToLower(game.GameType)}
	if extra, ok := byTitle[title]; ok {
		tags = append(tags, extra...)
	}

	return tags
}

func computeStats(games []CatalogGame) PageStats {
	stats := PageStats{EarliestYear: 9999}
	types := map[string]bool{}
	for _, game := range games {
		if game.ReleaseYear < stats.EarliestYear {
			stats.EarliestYear = game.ReleaseYear
		}
		if game.ReleaseYear > stats.LatestYear {
			stats.LatestYear = game.ReleaseYear
		}
		if strings.Contains(strings.ToLower(game.Developer), "bethesda") {
			stats.BethesdaGames++
		}
		stats.AverageRating += game.Rating
		types[game.GameType] = true
	}

	stats.TotalGames = len(games)
	stats.DistinctTypes = len(types)
	if stats.TotalGames > 0 {
		stats.AverageRating = stats.AverageRating / float64(stats.TotalGames)
	}
	if stats.EarliestYear == 9999 {
		stats.EarliestYear = 0
	}

	return stats
}

func filterGames(games []CatalogGame, tag, gameType, year, search string) []CatalogGame {
	filtered := make([]CatalogGame, 0, len(games))
	search = normalizeFilter(search)
	for _, game := range games {
		if tag != "" && !containsTag(game.Tags, tag) {
			continue
		}
		if gameType != "" && normalizeFilter(game.GameType) != gameType {
			continue
		}
		if year != "" && strconv.Itoa(game.ReleaseYear) != year {
			continue
		}
		if search != "" && !matchesSearch(game, search) {
			continue
		}
		filtered = append(filtered, game)
	}
	return filtered
}

func matchesSearch(game CatalogGame, search string) bool {
	fields := []string{game.Title, game.GameType, game.Description, game.Developer, game.Platforms, game.TagLine}
	for _, field := range fields {
		if strings.Contains(strings.ToLower(field), search) {
			return true
		}
	}
	return false
}

func containsTag(tags []string, target string) bool {
	target = normalizeFilter(target)
	for _, tag := range tags {
		if normalizeFilter(tag) == target {
			return true
		}
	}
	return false
}

func collectTags(games []CatalogGame) []string {
	set := map[string]bool{}
	tags := []string{}
	for _, game := range games {
		for _, tag := range game.Tags {
			if !set[tag] {
				set[tag] = true
				tags = append(tags, tag)
			}
		}
	}
	sort.Strings(tags)
	return tags
}

func collectTypes(games []CatalogGame) []string {
	set := map[string]bool{}
	types := []string{}
	for _, game := range games {
		typeName := normalizeFilter(game.GameType)
		if !set[typeName] {
			set[typeName] = true
			types = append(types, game.GameType)
		}
	}
	sort.Strings(types)
	return types
}

func collectYears(games []CatalogGame) []int {
	set := map[int]bool{}
	years := []int{}
	for _, game := range games {
		if !set[game.ReleaseYear] {
			set[game.ReleaseYear] = true
			years = append(years, game.ReleaseYear)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(years)))
	return years
}

func bestGames(games []CatalogGame, limit int) []CatalogGame {
	cloned := make([]CatalogGame, len(games))
	copy(cloned, games)
	sort.Slice(cloned, func(i, j int) bool {
		if cloned[i].Rating == cloned[j].Rating {
			return cloned[i].ReleaseYear > cloned[j].ReleaseYear
		}
		return cloned[i].Rating > cloned[j].Rating
	})
	if limit > len(cloned) {
		limit = len(cloned)
	}
	return cloned[:limit]
}

func findGameByID(games []CatalogGame, id int) (CatalogGame, bool) {
	for _, game := range games {
		if game.ID == id {
			return game, true
		}
	}
	return CatalogGame{}, false
}

func relatedGames(games []CatalogGame, current CatalogGame) []CatalogGame {
	pool := make([]CatalogGame, 0, len(games))
	for _, game := range games {
		if game.ID == current.ID {
			continue
		}
		pool = append(pool, game)
	}

	if len(pool) == 0 {
		return []CatalogGame{}
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(pool), func(i, j int) {
		pool[i], pool[j] = pool[j], pool[i]
	})

	result := make([]CatalogGame, 0, 3)
	for i := 0; i < len(pool) && i < 3; i++ {
		result = append(result, pool[i])
	}
	return result
}

func getIngameImages(gameID int) []string {
	gameFiles := map[int]string{
		1: "fallout1",
		2: "fallout2",
		3: "fallout3",
		4: "fallout_new_vegas",
		5: "fallout4",
		6: "fallout76",
		7: "fallout_shelter",
	}

	base, exists := gameFiles[gameID]
	if !exists {
		return []string{}
	}

	return []string{
		resolveAssetURL("/static/images_ingames/" + base + "_ingame1.jpg"),
		resolveAssetURL("/static/images_ingames/" + base + "_ingame2.jpg"),
		resolveAssetURL("/static/images_ingames/" + base + "_ingame3.jpg"),
	}
}

func resolveAssetURL(path string) string {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}

	baseURL := strings.TrimRight(strings.TrimSpace(os.Getenv("IMAGE_ASSET_BASE_URL")), "/")
	if baseURL == "" {
		return path
	}

	clean := strings.TrimPrefix(path, "/static/images/")
	clean = strings.TrimPrefix(clean, "/static/images_ingames/")
	clean = strings.TrimPrefix(clean, "/")
	return baseURL + "/" + clean
}

func normalizeFilter(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}
