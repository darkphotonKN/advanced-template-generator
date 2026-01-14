# DDD API Template Setup & Usage Guide

## Part 1: Setting Up the DDD Template (One-time Setup)

### Phase 1: Template Extraction from Cashflow

#### Step 1: Copy Base Project
```bash
# From go-template-generator directory
cp -r ../cashflow templates/ddd-api
cd templates/ddd-api
```

#### Step 2: Clean Project-Specific Files
```bash
# Remove unnecessary files
rm -rf .git
rm -rf tmp/
rm -rf bin/
rm -f go.sum
rm -rf .DS_Store
rm -f .env  # Keep .env.example

# Clean migrations except base schema
cd migrations
rm -f !(000001_init_schema.up.sql|000001_init_schema.down.sql)
cd ..
```

#### Step 3: Create Template File
```bash
# Rename go.mod to template
mv go.mod go.mod.tmpl
```

### Phase 2: Convert to Template Variables

#### File: `go.mod.tmpl`
```go
module {{.ModuleName}}

go 1.23.0

require (
    // ... dependencies remain the same
)
```

#### File: `docker-compose.yml`
Replace hardcoded values:
```yaml
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: {{.DBUser}}
      POSTGRES_PASSWORD: {{.DBPassword}}
      POSTGRES_DB: {{.DBName}}
    ports:
      - "{{.DBPort}}:5432"

  redis:
    image: redis:7-alpine
    ports:
      - "{{.RedisPort}}:6379"
```

#### File: `.env.example`
```env
# Database
DB_HOST=localhost
DB_PORT={{.DBPort}}
DB_USER={{.DBUser}}
DB_PASSWORD={{.DBPassword}}
DB_NAME={{.DBName}}

# Server
PORT={{.APIPort}}

# Redis
REDIS_HOST=localhost
REDIS_PORT={{.RedisPort}}

# JWT (conditional - only if auth enabled)
{{if .IncludeAuth}}JWT_SECRET=your-secret-key-here{{end}}

# AWS S3 (optional)
{{if .IncludeS3}}
AWS_REGION=us-east-1
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
S3_BUCKET=
{{end}}
```

#### File: `CLAUDE.md`
Update project description:
```markdown
## Project Overview

{{.ProjectDescription}}

## Quick Commands
...
```

### Phase 3: Entity Refactoring

#### Rename Domain Folder
```bash
# Template structure
internal/{{.PrimaryEntity}}/
├── model.go
├── repository.go
├── service.go
└── handler.go
```

#### Update `internal/{{.PrimaryEntity}}/model.go`
```go
package {{.PrimaryEntity}}

import (
    "time"
    "github.com/google/uuid"
)

type {{.EntityCapitalized}} struct {
    ID          uuid.UUID `json:"id" db:"id"`
    Name        string    `json:"name" db:"name"`
    Description string    `json:"description" db:"description"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Create{{.EntityCapitalized}}Request struct {
    Name        string `json:"name" binding:"required"`
    Description string `json:"description"`
}
```

#### Update `config/routes.go`
```go
package config

import (
    "database/sql"
    "log/slog"

    "github.com/gin-gonic/gin"
    "{{.ModuleName}}/internal/{{.PrimaryEntity}}"
    "{{.ModuleName}}/internal/middleware"
    {{if .IncludeS3}}"{{.ModuleName}}/internal/s3"{{end}}
)

func SetupRoutes(db *sql.DB, {{if .IncludeS3}}s3Service s3.Service, {{end}}logger *slog.Logger) *gin.Engine {
    router := gin.New()

    // Middleware
    router.Use(middleware.RequestID())
    router.Use(middleware.RequestLogger(logger))
    {{if .IncludeAuth}}router.Use(middleware.AuthMiddleware()){{end}}

    // Initialize services
    repo := {{.PrimaryEntity}}.NewRepository(db)
    service := {{.PrimaryEntity}}.NewService(repo, {{if .IncludeS3}}s3Service, {{end}}logger)
    handler := {{.PrimaryEntity}}.NewHandler(service, logger)

    // Routes
    api := router.Group("/api")
    {
        items := api.Group("/{{.EntityPlural}}")
        {
            items.POST("", handler.Create{{.EntityCapitalized}})
            items.GET("", handler.List{{.EntityPlural}})
            items.GET("/:id", handler.Get{{.EntityCapitalized}})
            items.PUT("/:id", handler.Update{{.EntityCapitalized}})
            items.DELETE("/:id", handler.Delete{{.EntityCapitalized}})
        }
    }

    return router
}
```

### Phase 4: Migration Templates

#### File: `migrations/000001_init_schema.up.sql`
```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS migrations_history (
    id SERIAL PRIMARY KEY,
    version INTEGER NOT NULL,
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### File: `migrations/000002_create_{{.PrimaryEntity}}_table.up.sql`
```sql
CREATE TABLE {{.EntityPlural}} (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_{{.EntityPlural}}_created_at ON {{.EntityPlural}}(created_at);
```

### Phase 5: Import Path Updates

#### Files to Update:
- All `.go` files in `cmd/`
- All `.go` files in `config/`
- All `.go` files in `internal/`

Replace:
```
"github.com/kranti/cashflow" → "{{.ModuleName}}"
"financial" → "{{.PrimaryEntity}}"
"Financial" → "{{.EntityCapitalized}}"
"transactions" → "{{.EntityPlural}}"
```

## Part 2: Using the Template to Generate Projects

### Prerequisites
```bash
# Build the generator tool (after template is ready)
cd generator
go build -o bin/go-gen cmd/main.go
```

### Basic Usage

#### Generate a Project with Auth
```bash
./bin/go-gen create my-app --entity=product
```

This will:
1. Create directory `my-app/`
2. Copy template files
3. Replace variables:
   - ProjectName: `my-app`
   - ModuleName: `github.com/kranti/my-app`
   - PrimaryEntity: `product`
   - EntityPlural: `products`
   - EntityCapitalized: `Product`
   - APIPort: `8010` (if first project)
   - DBPort: `5442`
   - RedisPort: `6389`
   - DBName: `my_app_db`
   - DBUser: `user`
   - DBPassword: `password`

#### Generate a Project without Auth
```bash
./bin/go-gen create my-app --entity=product --no-auth
```

### Post-Generation Steps (Automatic)
The generator will automatically:
```bash
cd my-app/
go mod init github.com/kranti/my-app
go mod tidy
rm -rf .git
git init
git add .
git commit -m "initial commit"
```

### Manual Setup (User runs these)
```bash
cd my-app/

# Copy env file
cp .env.example .env

# Start infrastructure
make docker-up

# Run migrations
make migrate-up

# Start development server
make dev
```

### Testing the Generated Project
```bash
# API should be running on allocated port
curl http://localhost:8010/health

# Create a product
curl -X POST http://localhost:8010/api/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Test Product", "description": "Test Description"}'

# List products
curl http://localhost:8010/api/products
```

## Generator Internals

### Port Allocation Algorithm
```go
func AllocatePorts(projectIndex int) {
    offset := projectIndex * 10
    return Ports{
        API:   8000 + offset,
        DB:    5432 + offset,
        Redis: 6379 + offset,
    }
}
```

### Project Registry (`~/.go-gen-projects.json`)
```json
{
  "projects": [
    {
      "name": "my-app",
      "index": 1,
      "api_port": 8010,
      "db_port": 5442,
      "redis_port": 6389,
      "entity": "product",
      "created_at": "2024-01-14T10:00:00Z"
    }
  ],
  "next_index": 2
}
```

### Template Variable Map
```go
vars := map[string]interface{}{
    "ProjectName":       "my-app",
    "ModuleName":        "github.com/kranti/my-app",
    "PrimaryEntity":     "product",
    "EntityPlural":      "products",
    "EntityCapitalized": "Product",
    "APIPort":           "8010",
    "DBPort":            "5442",
    "RedisPort":         "6389",
    "DBName":            "my_app_db",
    "DBUser":            "user",
    "DBPassword":        "password",
    "IncludeAuth":       true,
    "IncludeS3":         false,
    "ProjectDescription": "DDD API for product management",
}
```

## Troubleshooting

### Port Conflicts
If ports are already in use:
```bash
# Check what's using the port
lsof -i :8010

# Manually override in .env file
PORT=8090
```

### Database Connection Issues
```bash
# Check if containers are running
docker-compose ps

# Check logs
docker-compose logs postgres

# Restart containers
make docker-down
make docker-up
```

### Migration Failures
```bash
# Check migration status
make migrate-status

# Force specific version if needed
make migrate-force VERSION=1
```

## Next Steps
1. Complete template file conversions
2. Build generator CLI tool
3. Test with multiple projects
4. Add more customization options
5. Create additional templates (microservice, simple)