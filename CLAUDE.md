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
2. Process template variables (ProjectName, Entity, etc.)
3. Remove .tmpl extensions
4. Rename entity directories
5. Initialize Go module and git repository
6. Update project registry

### 6. Template Variables
Replace these in all .tmpl files:
- `{{.ProjectName}}` - User's project name (e.g., "todo-app")
- `{{.ModuleName}}` - github.com/kranti/{project-name}
- `{{.PrimaryEntity}}` - User's entity (e.g., "task")
- `{{.EntityCapitalized}}` - Capitalized entity (e.g., "Task")
- `{{.EntityPlural}}` - Plural form (e.g., "tasks")
- `{{.APIPort}}`, `{{.DBPort}}`, `{{.RedisPort}}` - Randomized ports
- `{{.IncludeAuth}}` - true/false for JWT authentication
- `{{.IncludeS3}}` - true/false for file uploads
- `{{.ProjectDescription}}` - User's description

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

**Persistent Settings:** `generator/config.yaml`
- Database credentials (user/password)
- Port allocation settings
- Default output location
- Template paths

**Project Registry:** `~/.go-gen-projects.json`
- Tracks all generated projects
- Manages port allocation
- Prevents naming conflicts

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