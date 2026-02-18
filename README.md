# 🎮 FALLOUT VAULT DATABASE

Base de données complète des jeux Fallout avec interface web style Pip-Boy.

## 📋 Description

Ce projet est une application web Go qui recense tous les jeux de la franchise Fallout avec un design rétro-futuriste inspiré du Pip-Boy. L'application utilise PostgreSQL pour stocker les données et peut être déployée gratuitement sur Scalingo.

## 🚀 Fonctionnalités

- **Page d'accueil** avec statistiques (nombre de jeux, années couvertes, note moyenne)
- **Catalogue complet** avec filtres par type de jeu et année
- **Pages détaillées** pour chaque jeu avec toutes les informations
- **Design post-apocalyptique** style Pip-Boy (vert phosphorescent #00FF41)
- **Responsive design** pour mobile et desktop
- **Effets visuels** : scanlines CRT, glow effects, animations

## 🛠️ Technologies

- **Backend** : Go 1.21+ avec framework Gin
- **Base de données** : PostgreSQL
- **Frontend** : Templates HTML + CSS + JavaScript vanilla
- **Déploiement** : Scalingo (compatible Heroku)

## 📂 Structure du projet

```
fallout-vault/
├── main.go                 # Point d'entrée du serveur
├── go.mod                  # Dépendances Go
├── Procfile               # Configuration Scalingo
├── database/
│   └── db.go              # Connexion PostgreSQL + tables
├── models/
│   └── game.go            # Modèle Game
├── handlers/
│   └── handlers.go        # Routes et logique
├── templates/
│   ├── base.html          # Template de base
│   ├── home.html          # Page d'accueil
│   ├── games.html         # Liste des jeux
│   └── game_detail.html   # Détail d'un jeu
└── static/
    ├── css/
    │   └── style.css      # Style Pip-Boy
    ├── js/
    │   └── main.js        # JavaScript interactif
    └── images/            # Images des jeux
```

## 🎮 Jeux inclus (données de test)

1. **Fallout** (1997) - RPG isométrique par Interplay
2. **Fallout 2** (1998) - RPG par Black Isle Studios
3. **Fallout 3** (2008) - Action RPG par Bethesda
4. **Fallout: New Vegas** (2010) - Action RPG par Obsidian
5. **Fallout 4** (2015) - Action RPG par Bethesda
6. **Fallout 76** (2018) - Online RPG par Bethesda
7. **Fallout Shelter** (2015) - Mobile/Strategy par Bethesda

## 🔧 Installation locale

### Prérequis

- Go 1.21 ou supérieur
- PostgreSQL 13 ou supérieur
- Git

### Étapes

1. **Cloner le projet** (ou créer les fichiers comme indiqué)

2. **Installer les dépendances Go**
   ```bash
   go mod download
   ```

3. **Configurer PostgreSQL**
   - Créer une base de données PostgreSQL
   - Définir la variable d'environnement `DATABASE_URL` :
   ```bash
   # Windows PowerShell
   $env:DATABASE_URL="postgresql://user:password@localhost:5432/fallout_vault?sslmode=disable"
   
   # Linux/Mac
   export DATABASE_URL="postgresql://user:password@localhost:5432/fallout_vault?sslmode=disable"
   ```

4. **Compiler et lancer le serveur**
   ```bash
   go build -o fallout-vault
   ./fallout-vault
   ```

5. **Accéder à l'application**
   - Ouvrir un navigateur sur `http://localhost:8080`

## 🌐 Déploiement sur Scalingo

### Configuration

1. **Créer un compte Scalingo** (30 jours gratuits)

2. **Installer Scalingo CLI**
   ```bash
   # Télécharger depuis https://cli.scalingo.com/
   ```

3. **Créer une application**
   ```bash
   scalingo create fallout-vault
   ```

4. **Ajouter PostgreSQL**
   ```bash
   scalingo --app fallout-vault addons-add postgresql postgresql-starter-512
   ```
   La variable `DATABASE_URL` sera automatiquement configurée.

5. **Déployer**
   ```bash
   git init
   git add .
   git commit -m "Initial commit"
   git push scalingo master
   ```

6. **Ouvrir l'application**
   ```bash
   scalingo --app fallout-vault open
   ```

### Variables d'environnement automatiques

Scalingo configure automatiquement :
- `DATABASE_URL` : URL de connexion PostgreSQL
- `PORT` : Port sur lequel le serveur doit écouter

## 🎨 Design & Palette de couleurs

- **Vert Pip-Boy** : #00FF41 (texte principal)
- **Noir** : #0A0E0D (fond)
- **Gris foncé** : #1A1D1A (cartes/zones)
- **Orange rouille** : #C8511B (accents/badges)
- **Jaune warning** : #FFD700 (titres importants)

### Effets visuels
- Scanlines CRT animées
- Text glow sur le vert phosphorescent
- Animations de pulse et blink
- Hover effects avec box-shadow
- Responsive grid layout

## 📡 Routes disponibles

- `GET /` - Page d'accueil avec statistiques
- `GET /games` - Liste de tous les jeux
- `GET /games?type=RPG` - Filtrer par type
- `GET /games?year=2010` - Filtrer par année
- `GET /games/:id` - Détail d'un jeu spécifique

## 🔐 Base de données

### Table `games`

| Colonne | Type | Description |
|---------|------|-------------|
| id | SERIAL | Identifiant unique |
| title | VARCHAR(255) | Nom du jeu |
| game_type | VARCHAR(100) | Type (RPG, Action RPG, etc.) |
| description | TEXT | Description du jeu |
| release_year | INTEGER | Année de sortie |
| developer | VARCHAR(255) | Studio développeur |
| platforms | VARCHAR(255) | Plateformes disponibles |
| rating | DECIMAL(3,1) | Note sur 10 |
| image_url | VARCHAR(500) | URL de l'image |
| created_at | TIMESTAMP | Date d'ajout |

## 🎯 Easter Eggs

- **Konami Code** : Tapez ↑↑↓↓←→←→BA pour activer le mode Vault-Boy
- Effets sonores dans la console navigateur
- Messages cachés dans les logs

## 📝 TODO / Améliorations futures

- [ ] Système d'authentification utilisateur
- [ ] Possibilité d'ajouter/modifier des jeux via interface
- [ ] Système de favoris/wishlist
- [ ] Commentaires et reviews utilisateurs
- [ ] API REST JSON pour développeurs
- [ ] Upload d'images directement sur le serveur
- [ ] Recherche en full-text
- [ ] Pagination pour grandes collections
- [ ] Export en PDF/JSON

## 🐛 Dépannage

### Erreur de connexion à la base de données
- Vérifier que `DATABASE_URL` est bien définie
- Vérifier que PostgreSQL est démarré
- Vérifier les credentials dans l'URL

### Le serveur ne démarre pas
- Vérifier que le port 8080 n'est pas déjà utilisé
- Compiler avec `go build` pour voir les erreurs

### Les templates ne se chargent pas
- Vérifier que le dossier `templates/` existe
- Vérifier les chemins relatifs dans `main.go`

## 📄 Licence

Ce projet est sous licence MIT. Libre d'utilisation et de modification.

## 👤 Auteur

Créé avec ❤️ pour les fans de Fallout

---

**⚡ WAR. WAR NEVER CHANGES. ⚡**
