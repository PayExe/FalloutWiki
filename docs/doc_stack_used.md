# Infrastructure et API

## Supabase

- Supabase sert de base de donnees PostgreSQL.
- L'application se connecte avec la variable `DATABASE_URL`.
- Si `DATABASE_URL` est vide, le serveur ne demarre pas.
- Les identifiants de connexion restent dans les variables d'environnement, pas dans le code.
- elle me sert aussi lors de la création de l'api publique avec un bucket ou je vais y stocker toute mes images

## Render

- Render heberge l'application web.
- Le service lance le projet avec le `Procfile`.
- Render fournit la variable `PORT` au demarrage.
- Il faut aussi configurer `DATABASE_URL`, `ADMIN_CODE` et, si besoin, `IMAGE_ASSET_BASE_URL` (ils sont dans le .env local mais surtout sur render)

## Base de donnees

- La base contient une table principale : `games`.
- Chaque jeu a un `ID`, un `Title`, un `GameType`, une `Description`, une `ReleaseYear`, un `Developer`, des `Platforms`, une `Rating`, une `ImageURL` et des `Tags`.
- Le code cree la structure avec GORM au demarrage.
- Si la table est vide, le projet ajoute des donnees de depart avec le seed.

## Routes / API du projet

- Ce projet n'expose pas une API JSON publique classique.
- Il expose surtout des routes web pour afficher les pages.

### Pages publiques

- `GET /` : page d'accueil.
- `GET /games` : catalogue des jeux.
- `GET /games/:id` : fiche detaillee d'un jeu.
- `GET /credits` : page des credits.

### Administration

- `GET /admin` : page de connexion admin.
- `POST /admin` : verification du code admin.
- `GET /admin/logout` : deconnexion.
- `GET /admin/games` : liste admin des jeux.
- `GET /admin/games/new` : formulaire de creation.
- `POST /admin/games` : creation d'un jeu.
- `GET /admin/games/:id/edit` : formulaire de modification.
- `POST /admin/games/:id/update` : mise a jour.
- `POST /admin/games/:id/delete` : suppression.
- `POST /admin/games/clear-tags` : suppression des tags.

## Variables importantes

- `DATABASE_URL` : connexion a Supabase.
- `PORT` : port utilise par Render ou en local.
- `ADMIN_CODE` : code pour acceder a l'admin.
- `IMAGE_ASSET_BASE_URL` : chemin ou domaine optionnel pour les images.
