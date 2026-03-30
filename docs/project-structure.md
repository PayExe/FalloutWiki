# Project Structure

## Root
- `main.go`: starts the server and registers routes.
- `Procfile`: deployment start command.
- `.env.example`: environment variable template.

## App code
- `database/`: PostgreSQL and seed logic.
- `handlers/`: route handlers and page data.
- `models/`: GORM models.

## Frontend
- `templates/`: HTML pages.
- `static/css/`: styles.
- `static/js/`: small front-end behavior.

## Notes
- Keep build outputs out of the repo.
- Keep secrets only in `.env` or hosting env vars.
