# Fallout Wiki

Projet scolaire qui à pour but de représenté un catalogue simple pour parcourir toute la licence Fallout, avec des fiches détaillées, des captures d'écran, un CRUD fonctionnel et protégé déployé directement sur render et en utilisant Supabase

## Lancer le projet

```bash
# 1. Cloner le repo
git clone https://github.com/ton-user/fallout-wiki.git
cd fallout-wiki

# 2. Créer le fichier .env
cp .env.example .env

# 3. Remplir les variables
DATABASE_URL=postgresql://...
ADMIN_CODE=ton-code
PORT=8080

# 4. Lancer
go run main.go
```

Ouvre `http://localhost:8080`.

## Déploiement

https://falloutwiki.onrender.com/

## Stack

- **Go 1.24** — langage principal
- **Gin** — serveur HTTP
- **GORM** — ORM pour la base de données
- **PostgreSQL (Supabase)** — base de données
- **HTML / CSS / JS** — frontend sans framework
- **Render** — hébergement

## Structure du projet

```
.
├── main.go                  # point d'entrée, routes
├── Procfile                 # commande de démarrage Render
├── .env.example             # modèle de variables d'env
├── go.mod / go.sum          # dépendances Go
│
├── database/
│   └── db.go                # connexion PostgreSQL + seed
│
├── models/
│   └── game.go              # modèle GORM (table games)
│
├── handlers/
│   └── handlers.go          # logique des routes
│
├── templates/
│   ├── base.html            # layout principal
│   ├── home.html            # accueil
│   ├── games.html           # catalogue
│   ├── game_detail.html     # fiche d'un jeu
│   ├── credits.html         # crédits
│   ├── admin_login.html     # connexion admin
│   ├── admin_games.html     # liste admin
│   └── admin_game_form.html # formulaire admin
│
├── static/
│   ├── css/style.css        # styles
│   └── js/main.js           # lightbox + animations
│
└── docs/
    ├── project-structure.md
    └── infra-et-api.md
```

## Supabase

- Fournit la base de données PostgreSQL
- On copie la chaîne de connexion dans `DATABASE_URL`
- Optionnel : créer un bucket Storage public `images` pour y stocker les visuels
- Les images s'appellent `fallout1_ingame1.jpg`, `fallout4_ingame2.jpg`, etc.

## Render

- Héberge l'application en ligne
- Lit le `Procfile` (`web: fallout-vault`)
- Variables à configurer : `DATABASE_URL`, `ADMIN_CODE`, `PORT`, `IMAGE_ASSET_BASE_URL`
- Déploiement automatique à chaque push sur la branche principale
