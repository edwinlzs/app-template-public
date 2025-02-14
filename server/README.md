# Dev Setup

```sh

```

# `.env` file

```sh
# AUTH
JWT_SECRET=<JWT_SECRET>

# DB
PGUSER=...
PGPASSWORD=...
PGDATABASE=...
PGHOST=localhost
PGPORT=5432

# DB MIGRATIONS
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://<PGUSER>:<PGPASSWORD>@<PGHOST>:<PGPORT>/<PGDATABASE>?sslmode=disable
GOOSE_MIGRATION_DIR=./sql/migrations
```

# GORM

- Pointers (`*string`) indicate nullable fields
- GORM auto handles `CreatedAt` and `UpdatedAt` fields
- embed `gorm.Model` to include commonly used fields
- ID is default primary key for a model
- `User` becomes `users` table in db

# Testing

- Run specific test `go test <module> -run <Testname regexp>`

# Resources

- Passing Server Env to Handlers - https://blog.questionable.services/article/http-handler-error-handling-revisited/
- Writing Middleware - https://blog.questionable.services/article/guide-logging-middleware-go/
- Golang JWT with supabase - https://depshub.com/blog/using-supabase-auth-as-a-service-with-a-custom-backend/
- Passing data from middleware to handler - https://stackoverflow.com/questions/75474701/passing-data-from-handler-to-middleware-after-serving-request
