# Go Template Generator

Generates production-ready Go DDD API projects with clean architecture, hot reload, Docker setup, and optional Next.js frontend.

## Features

### Backend (Go)
- **DDD Architecture**: Domain-Driven Design with clean separation
- **Hot Reload**: Pre-configured Air for development
- **Docker Ready**: PostgreSQL, Redis with Docker Compose
- **Database Migrations**: Built-in migration support
- **Optional Auth**: JWT authentication (configurable)
- **SOLID Principles**: ISP, DIP, and clean separation

### Frontend (Next.js) - Optional
- **Next.js 15**: Latest React framework with App Router
- **TypeScript**: Full type safety
- **Tailwind CSS**: Utility-first styling with shadcn/ui components
- **TanStack Query**: Data fetching and caching
- **Zustand**: State management
- **Axios**: HTTP client with interceptors
- **Full CRUD UI**: Pre-built components for entity management

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

### Quick Setup
Most users only need to change one setting in `generator/config.yaml`:
```yaml
defaults:
  module_prefix: "github.com/yourusername/"  # Change to your GitHub username
```

### What's Preset vs What's Prompted

**Preset in config.yaml (applies to all projects):**
- Database: `user`/`password` credentials
- Ports: Base values with ±50 randomization
- Location: Projects created as siblings (`../`)
- Module: `github.com/kranti/` prefix

**Claude prompts for (per project):**
- Project name
- Primary entity (e.g., task, product, post)
- Authentication needed? (default: yes)
- S3 uploads needed? (default: no)
- Frontend needed? (default: no)
- Project description

### Project Location
By default, projects are generated in the parent directory:
```
/your-go-projects/
├── go-template-generator/    # This repository
├── todo-app/                # Generated project
├── inventory-api/           # Generated project
└── blog-api/               # Generated project
```

### Advanced Configuration
See [Configuration Guide](docs/CONFIG_GUIDE.md) for detailed options including:
- Custom port ranges
- Different database credentials
- Alternative output locations
- Feature flag defaults

## Usage Examples

### Simple Todo App (No Auth)
**Prompt:** "Create a todo app API without authentication"
**Result:** `../todo-app/` with task management endpoints

### Full-Stack App with Frontend
**Prompt:** "Create a task management app with frontend and authentication"
**Result:**
- `../task-manager/` - Go API
- `../task-manager-frontend/` - Next.js app

### E-commerce with Auth
**Prompt:** "Generate an e-commerce product API with JWT authentication"
**Result:** `../ecommerce-api/` with protected product endpoints

### Blog Platform with UI
**Prompt:** "Create a blog platform with frontend, authentication, and image uploads"
**Result:**
- `../blog-platform/` - Go API with S3 support
- `../blog-platform-frontend/` - Full React UI

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
