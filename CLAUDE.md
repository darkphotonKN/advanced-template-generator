# Claude Code Instructions for Go Template Generator

This project generates production-ready Go DDD API projects with clean architecture, hot reload, and Docker setup.

## How to Generate a New Project

When the user requests a new Go API project, follow these steps:

### 1. Start in This Directory
Always ensure you're in the go-template-generator root directory when starting.

### 2. Navigate to Generator
```bash
cd generator
```

### 3. Follow Interactive Workflow
Use the structured workflow from `docs/CLAUDE_GENERATION_WORKFLOW.md` to:
- Gather project requirements (name, entity, auth, S3, description)
- Validate inputs using `docs/VALIDATION_CHECKLIST.md`
- Use prompts from `prompts/PROJECT_GENERATION_PROMPTS.md`

### 4. Project Location
Generated projects are created as siblings to this directory:
```
/Users/kranti/Documents/Code/Go/
├── go-template-generator/    # This repository
├── todo-app/                # Generated project
├── inventory-api/           # Another generated project
└── blog-api/               # Another generated project
```

Default output location is configured in `generator/config.yaml` as `../`

### 5. Generation Process
1. Copy templates from `templates/ddd-api/`
2. Process only the 3 configuration template files (.tmpl)
3. Update module name in go.mod and all imports
4. Initialize git repository
5. Update project registry

### 6. Template Information
The template provides a working example API with:
- **Example Entity**: "item" - a simple CRUD example
- **Clean Architecture**: Repository, Service, Handler layers
- **Error Handling**: Professional error utilities included
- **No Template Variables in Code**: All Go code is ready to run

Only configuration files need processing:
- `docker-compose.yml.tmpl` - for port configuration
- `.env.example.tmpl` - for environment setup
- Module name updates in go.mod and imports

## Example User Prompts

**Simple Request:**
"Create a todo app API without authentication"

**Detailed Request:**
"Generate an e-commerce product management API with JWT auth and S3 image uploads"

**Your Response Pattern:**
1. Start with greeting from `prompts/PROJECT_GENERATION_PROMPTS.md`
2. Ask for project name, entity, auth, S3, description
3. Validate inputs and confirm configuration
4. Execute generation process
5. Provide success message with next steps

## Configuration

### What's Preset (You Don't Need to Ask)
These are configured in `generator/config.yaml` and apply to ALL projects:

**Fixed Settings:**
- Database credentials: Always `user`/`password`
- Module prefix: `github.com/kranti/`
- Output location: `../` (creates sibling directories)
- Port base values: API 8000, DB 5432, Redis 6379
- Port randomization: ±50 range for each service
- Git initialization: Always enabled

**Default Features (but still ask to confirm):**
- Authentication: `true` by default
- S3 uploads: `false` by default
- Redis caching: Always included (never ask)

### What Claude Prompts For (Per Project)
You MUST gather these for each project:

1. **Project Name** - Required (can infer from request)
2. **Primary Entity** - Required (must ask if not clear)
3. **Authentication** - Ask yes/no (default: yes)
4. **S3 Uploads** - Ask yes/no (default: no)
5. **Description** - Required for documentation

### How to Modify Defaults
Users can edit `generator/config.yaml` if they need:
- Different GitHub username (change `module_prefix`)
- Different database credentials
- Different port ranges
- Different output location

**Project Registry:** `~/.go-gen-projects.json`
- Automatically managed
- Tracks port allocations
- Prevents conflicts

## Documentation Reference

- `docs/CLAUDE_GENERATION_WORKFLOW.md` - Complete step-by-step workflow
- `prompts/PROJECT_GENERATION_PROMPTS.md` - Standardized user prompts
- `docs/VALIDATION_CHECKLIST.md` - Input validation rules
- `examples/EXAMPLE_CONVERSATIONS.md` - Example interactions
- `docs/TROUBLESHOOTING_GUIDE.md` - Error resolution

## Quick Start Example

User: "Create a task management API without auth"

Your workflow:
1. Ask for project name → "task-manager"
2. Confirm entity → "task"
3. Confirm no auth → "no"
4. Confirm no S3 → "no"
5. Get description → "Simple task management API"
6. Generate at `../task-manager/`
7. Provide setup instructions

The generated project will be ready to run with:
```bash
cd ../task-manager
cp .env.example .env
make docker-up && make migrate-up && make dev
```

Always reference the detailed documentation files for validation rules, error handling, and troubleshooting procedures.