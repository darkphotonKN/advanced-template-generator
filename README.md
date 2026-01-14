# Go Template Generator

Generates production-ready Go DDD API projects with clean architecture, hot reload, and Docker setup.

## Features

- **DDD Architecture**: Domain-Driven Design with clean separation
- **Hot Reload**: Pre-configured Air for development
- **Docker Ready**: PostgreSQL, Redis with Docker Compose
- **Database Migrations**: Built-in migration support
- **Optional Auth**: JWT authentication (configurable)
- **Zero Config Start**: Working project in seconds
- **Claude Integration**: AI-assisted project generation
- **SOLID Principles**: ISP, DIP, and clean separation

## Setup for Claude Code

### 1. Clone and Setup
```bash
git clone https://github.com/kranti/go-template-generator.git
cd go-template-generator
```

### 2. Using with Claude Code
Prompt Claude from this directory:
```
Create a todo app API without authentication
```

Claude will:
1. Navigate to the `generator/` directory
2. Follow interactive prompts for project configuration
3. Generate the project in a sibling directory (e.g., `../todo-app/`)
4. Provide setup instructions

### 3. Start Your Generated Project
```bash
cd ../your-project-name/
cp .env.example .env
make docker-up && make migrate-up && make dev
```

## Configuration

### Persistent Settings
Edit `generator/config.yaml` to modify:
- **Output Location**: `output_location: "../"` (default: creates projects as siblings)
- **Database Credentials**: Always `user`/`password`
- **Port Allocation**: Base ports and randomization settings
- **Default Features**: Auth, S3, Redis configuration

### Project Location
By default, projects are generated in the parent directory:
```
/your-go-projects/
├── go-template-generator/    # This repository
├── todo-app/                # Generated project
├── inventory-api/           # Generated project
└── blog-api/               # Generated project
```

## Usage Examples

### Simple Todo App (No Auth)
**Prompt:** "Create a todo app API without authentication"
**Result:** `../todo-app/` with task management endpoints

### E-commerce with Auth
**Prompt:** "Generate an e-commerce product API with JWT authentication"
**Result:** `../ecommerce-api/` with protected product endpoints

### Blog Platform
**Prompt:** "Create a blog API with post management and authentication"
**Result:** `../blog-api/` with post CRUD and JWT middleware

### Generated Project Structure

```
my-project/
├── cmd/
│   └── main.go              # Application entry point
├── config/
│   ├── database.go          # Database configuration
│   └── routes.go            # Route definitions
├── internal/
│   ├── {entity}/            # Your domain (product, post, etc.)
│   │   ├── model.go         # Domain models
│   │   ├── repository.go    # Data access layer
│   │   ├── service.go       # Business logic
│   │   └── handler.go       # HTTP handlers
│   ├── middleware/          # HTTP middleware
│   └── util/                # Utility functions
├── migrations/              # Database migrations
├── .air.toml                # Hot reload config
├── docker-compose.yml       # Infrastructure setup
├── Makefile                 # Development commands
└── CLAUDE.md                # AI assistant guide
```

## Dynamic Port Allocation

Each project gets unique randomized ports to avoid conflicts:

| Project | API Port | DB Port  | Redis Port |
| ------- | -------- | -------- | ---------- |
| 1st     | ~8010±50 | ~5442±50 | ~6389±50   |
| 2nd     | ~8020±50 | ~5452±50 | ~6399±50   |
| 3rd     | ~8030±50 | ~5462±50 | ~6409±50   |

**Example**: Project 1 might get API:8030, DB:5473, Redis:6435 instead of predictable ports.
**Safety**: All ports guaranteed within valid system range (1024-65535).

## Database Configuration

All projects use consistent credentials:

- **User**: `user`
- **Password**: `password`
- **Database**: `{project_name}_db`

## Available Commands

After generation, each project includes:

```bash
make dev          # Run with hot reload
make build        # Build binary
make test         # Run tests
make lint         # Run linter
make docker-up    # Start PostgreSQL & Redis
make docker-down  # Stop containers
make migrate-up   # Run migrations
make migrate-down # Rollback migrations
```

## Git Workflow

Projects are initialized with git and follow conventional commits:

```bash
git commit -m "feat: add user authentication"
git commit -m "fix: resolve database timeout"
git commit -m "refactor: simplify service layer"
git commit -m "docs: update API documentation"
```

## Documentation

- **CLAUDE.md** - Primary instructions for Claude Code
- **docs/CLAUDE_GENERATION_WORKFLOW.md** - Complete generation workflow
- **prompts/PROJECT_GENERATION_PROMPTS.md** - Standardized prompts
- **docs/VALIDATION_CHECKLIST.md** - Input validation rules
- **examples/EXAMPLE_CONVERSATIONS.md** - Example interactions
- **docs/TROUBLESHOOTING_GUIDE.md** - Error resolution

## Project Registry

The generator tracks all projects in `~/.go-gen-projects.json` for port allocation and conflict prevention.

## Architecture

### Generated Project Structure
```
my-project/
├── cmd/main.go              # Application entry point
├── internal/{entity}/       # Domain layer (task, product, etc.)
│   ├── model.go             # Domain models
│   ├── repository.go        # Data access
│   ├── service.go           # Business logic
│   └── handler.go           # HTTP handlers
├── config/routes.go         # API routes
├── migrations/              # Database schema
├── docker-compose.yml       # Infrastructure
└── CLAUDE.md               # Development guide
```

### Principles
- **SOLID Principles**: Single Responsibility, Interface Segregation, Dependency Inversion
- **Clean Architecture**: Separation of concerns, testable code
- **12-Factor App**: Environment-based configuration
- **Domain-Driven Design**: Business logic in domain layer

## Contributing

Contributions are welcome! The template is based on the successful `cashflow` project structure.

## License

MIT
