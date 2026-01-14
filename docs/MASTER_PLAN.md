# Go Template Generator - Master Plan

## Overview
A comprehensive template generation system for Go projects, starting with DDD API templates and expanding to support multiple project types in the future.

## Phase 1: DDD API Generator (Current)
Focus on creating a robust Domain-Driven Design API generator based on the successful cashflow project structure.

### Goals
- Extract and templatize the cashflow project
- Create consistent port and database naming conventions
- Support auth/no-auth configurations
- Maintain SOLID principles and clean architecture

### Key Features
- Automatic port allocation (prevents conflicts)
- Standardized database credentials
- Git initialization with conventional commits
- Hot reload development setup
- Docker Compose orchestration
- Database migrations

## Phase 2: Simple/Script Generator (Future)
- CLI tools with Cobra
- Simple scripts without complex architecture
- Minimal dependencies
- Quick prototyping support

## Phase 3: Microservice Generator (Future)
Based on cosmic-void project patterns:
- Go workspace management
- gRPC service communication
- Service discovery with Consul
- Message broker integration (RabbitMQ/Kafka)
- API Gateway pattern
- Common shared modules

## Technical Decisions

### Fixed Conventions
- Database user: `user`
- Database password: `password`
- Database naming: `{project}_db`
- Port increment: 10 per project
- Module prefix: `github.com/kranti/`

### Port Allocation Strategy
```
Project 1: API 8000, DB 5432, Redis 6379
Project 2: API 8010, DB 5442, Redis 6389
Project 3: API 8020, DB 5452, Redis 6399
```

### Git Commit Convention
Following Conventional Commits specification:
- `feat`: New feature
- `fix`: Bug fix
- `refactor`: Code refactoring
- `docs`: Documentation
- `test`: Testing
- `chore`: Maintenance
- `style`: Formatting
- `perf`: Performance
- `ci`: CI/CD changes
- `build`: Build system

## Directory Structure
```
go-template-generator/
├── docs/
│   ├── MASTER_PLAN.md           # This file
│   └── DDD_API_SETUP.md         # DDD API implementation guide
├── templates/
│   ├── ddd-api/                 # DDD API template
│   ├── simple/                  # (Future) Simple script template
│   └── microservice/            # (Future) Microservice template
├── generator/
│   ├── cmd/                     # CLI entry point
│   ├── internal/                # Core logic
│   └── config.yaml              # Permanent settings
└── scripts/                     # Helper scripts
```

## Project Registry
Maintains `~/.go-gen-projects.json` to track:
- Generated projects
- Allocated ports
- Creation timestamps
- Project indices

## Success Metrics
- [ ] Generate DDD API project in < 30 seconds
- [ ] Zero port conflicts between projects
- [ ] Consistent structure across all generated projects
- [ ] Working hot reload out of the box
- [ ] Database migrations run successfully
- [ ] All tests pass in generated projects

## Next Steps
1. Complete DDD API template extraction
2. Build generator CLI
3. Test with 3+ sample projects
4. Document customization points
5. Plan Phase 2 (Simple templates)