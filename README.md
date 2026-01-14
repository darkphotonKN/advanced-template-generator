# Go Template Generator

A powerful Go project generator that creates production-ready applications with consistent structure, best practices, and pre-configured development environment.

## Features

- **DDD Architecture**: Domain-Driven Design with clean architecture principles
- **Hot Reload**: Pre-configured Air for development
- **Docker Ready**: PostgreSQL, Redis with Docker Compose
- **Database Migrations**: Built-in migration support
- **Optional Auth**: JWT authentication (configurable)
- **Zero Config Start**: Working project in seconds
- **AI-Ready**: CLAUDE.md included for AI assistance
- **SOLID Principles**: ISP, DIP, and clean separation

## Quick Start

```bash
# Generate a new project
./generator/bin/go-gen create my-app --entity=product

# Navigate to project
cd my-app/

# Start infrastructure
make docker-up

# Run migrations
make migrate-up

# Start development server
make dev
```

Your API is now running at `http://localhost:8010` 🎉

## Installation

```bash
# Clone the repository
git clone https://github.com/kranti/go-template-generator.git
cd go-template-generator

# Build the generator
cd generator
go build -o bin/go-gen cmd/main.go
```

## Usage

### Basic Project Generation

```bash
# With authentication (default)
./bin/go-gen create my-project --entity=item

# Without authentication
./bin/go-gen create my-project --entity=item --no-auth

# Custom entity
./bin/go-gen create blog --entity=post
```

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

## Dynamic Port Allocation 🎲

Each project gets unique randomized ports to avoid conflicts and prevent predictable patterns:

| Project | API Port     | DB Port      | Redis Port   |
|---------|--------------|--------------|--------------|
| 1st     | ~8010±50     | ~5442±50     | ~6389±50     |
| 2nd     | ~8020±50     | ~5452±50     | ~6399±50     |
| 3rd     | ~8030±50     | ~5462±50     | ~6409±50     |

**Example**: Project 1 might get API:8030, DB:5473, Redis:6435 instead of the predictable 8010, 5442, 6389.
**Safety**: All ports guaranteed to be within valid system range (1024-65535).

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

## Architecture Principles

Generated projects follow:

- **SOLID Principles**: Single Responsibility, Interface Segregation, Dependency Inversion
- **Clean Architecture**: Separation of concerns, testable code
- **12-Factor App**: Environment-based configuration
- **Domain-Driven Design**: Business logic in domain layer

## Project Registry

The generator tracks all created projects in `~/.go-gen-projects.json` to manage port allocation and prevent conflicts.

## Documentation

- [Master Plan](docs/MASTER_PLAN.md) - Overall system design and roadmap
- [DDD API Setup](docs/DDD_API_SETUP.md) - Detailed setup and usage guide

## Examples

### E-commerce API

```bash
./bin/go-gen create ecommerce --entity=product
# Creates product management API with auth
```

### Task Manager

```bash
./bin/go-gen create tasks --entity=task --no-auth
# Creates task API without authentication
```

### Blog Platform

```bash
./bin/go-gen create blog --entity=post
# Creates blog API with post management
```

## Contributing

Contributions are welcome! The template is based on the successful `cashflow` project structure.

## Roadmap

- [x] DDD API Template
- [ ] Simple Script Template
- [ ] Microservice Template
- [ ] GraphQL API Template
- [ ] CLI Tool Template

## License

MIT

## Support

For issues or questions, please open an issue on GitHub.

---

Built with ❤️ for the Go community
