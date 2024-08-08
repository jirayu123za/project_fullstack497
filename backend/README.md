## Command

```bash
go run cmd/main.go

nodemon --exec go run cmd/main.go --signal SIGTERM
```

## .ENV

```bash
PORT = 3000

POSTGRES_PASSWORD=fullstack497.
POSTGRES_USER=fullstack497
POSTGRES_DB=g2_db
POSTGRES_PORT=5432
POSTGRES_HOST=localhost

POSTGRES_APP_USER=appuser
POSTGRES_APP_PASSWORD=1234

TIMEZONE=Asia/Bangkok

DATABASE_DSN = host = localhost user = fullstack497 password = fullstack497. dbname = g2_db port = 5432 sslmode = disable TimeZone = Asia/Bangkok
#DATABASE_DSN=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable&timezone=${TIMEZONE}

# change host from localhost to db
```
