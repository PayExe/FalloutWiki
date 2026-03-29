# FalloutWiki Experimental Project

This is an experimental database project about the Fallout game series, created and deployed by a French Fallout fan. The goal is to explore web development, database integration, and deployment using modern technologies.

## Features
- Game catalogue for the Fallout franchise
- Game details, descriptions, images, and stats
- Simple web interface (HTML/CSS/JS)
- PostgreSQL database (via GORM)
- Go backend (Gin framework)

## How It Works
- `models.Game` describes one Fallout game in the database.
- GORM manages the table automatically with `AutoMigrate`.
- The public site reads the data and shows it in pages.
- The admin pages let you create, edit, and delete games.

## CRUD On `games`
- Create: add a new game from `/admin/games/new`
- Read: display the catalogue and the detail page
- Update: edit an existing game from the admin list
- Delete: remove a game from the admin list

## Simple Explanation For Oral Defense
- I used Go + Gin for the server because it is clean and easy to route.
- I used GORM so the database table is managed by an ORM, not raw SQL.
- I used PostgreSQL to store the games data permanently.
- I added an admin CRUD so the project matches a real database use case.

## Technologies Used
- Go (Gin, GORM)
- PostgreSQL
- HTML/CSS/JS

## Deployment
The project is designed for learning and experimentation. It can be deployed locally or on platforms like Scalingo.

## Setup
1. Clone the repository
2. Copy `.env.example` to `.env` and set your PostgreSQL `DATABASE_URL`
3. Optional: set `PORT` if your environment requires it
4. Run the server:
    ```bash
    go run main.go
    ```

## Environment Variables
- `DATABASE_URL`: required, PostgreSQL connection string
- `PORT`: optional, defaults to `8080`

## Database Notes
- The project uses one main table: `games`.
- The seed data fills the table only if it is empty.
- This keeps the project easy to test locally.

## Disclaimer
This project is not official and is for educational purposes only. All Fallout content belongs to their respective owners.

---
Made by a French Fallout player, for fun and learning.
