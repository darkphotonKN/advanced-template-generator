# Claude Code Instructions for Go Template Generator

This project generates production-ready Go DDD API projects with clean architecture, hot reload, Docker setup, and optional Next.js frontend.

## Reference Docs for This Generator

Before implementing features or making changes to the generator itself, check the relevant doc in `docs/`:

| Task | Read First |
|------|------------|
| Generation workflow | `docs/CLAUDE_GENERATION_WORKFLOW.md` |
| Validation rules | `docs/VALIDATION_CHECKLIST.md` |
| Configuration options | `docs/CONFIG_GUIDE.md` |
| Frontend architecture | `docs/FRONTEND_ARCHITECTURE.md` |
| Troubleshooting | `docs/TROUBLESHOOTING_GUIDE.md` |
| Example conversations | `examples/EXAMPLE_CONVERSATIONS.md` |
| Generation prompts | `prompts/PROJECT_GENERATION_PROMPTS.md` |
| Generated plans | `docs/plans/*.md` |

**Do NOT guess configuration options, validation rules, or workflow steps.** The docs have the exact procedures to follow.

The docs/ folder contains detailed specifications and procedures extracted from design documents and requirements.

If a doc doesn't exist for what you're implementing, create it first or ask for clarification. If there are very sensible defaults for a configuration or workflow then you can use that while asking for confirmation.

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
- Gather project requirements (name, entity, auth, S3, frontend, description)
- Validate inputs using `docs/VALIDATION_CHECKLIST.md`
- Use prompts from `prompts/PROJECT_GENERATION_PROMPTS.md`

### 4. Project Location
Generated projects are created as siblings to this directory:

**Backend Only:**
```
/Users/kranti/Documents/Code/Go/
├── go-template-generator/    # This repository
├── todo-app/                # Generated API project (backend only)
├── inventory-api/           # Another API project (backend only)
└── blog-api/               # Another API project (backend only)
```

**Full-Stack (Backend + Frontend):**
```
/Users/kranti/Documents/Code/Go/
├── go-template-generator/    # This repository
└── todo-app/                 # App container folder
    ├── todo-app-server/      # Go API backend
    └── todo-app-client/      # Next.js frontend
```

Default output location is configured in `generator/config.yaml` as `../`

### 5. Generation Process

**For Backend (Always Generated):**
1. Copy templates from `templates/ddd-api/`
2. Process template files (.tmpl) including SPECIFICATION.md.tmpl with project-specific details
3. Update module name in go.mod and all import paths automatically
4. Rename "item" entity to user's chosen entity throughout codebase
5. Generate SPECIFICATION.md with project requirements and features
6. Include AGENTS.md for test agent support
7. Initialize git repository with correct module name
8. Update project registry

**For Frontend (If Requested):**
1. Use the provided frontend generation script:
   ```bash
   # From generator/ directory:
   ./scripts/generate-frontend.sh project-name entity api-port frontend-port [auth] [s3]

   # Example:
   ./scripts/generate-frontend.sh todo-app task 8050 3050 true false
   ```

   This script will:
   - Copy templates from `templates/nextjs-frontend/`
   - Process configuration template files (.tmpl)
   - Rename "item" entity to user's chosen entity throughout
   - Update import paths and component references
   - Generate SPECIFICATION.md with frontend requirements
   - Initialize separate git repository
   - Provide setup instructions (but NOT run npm install automatically)

### 6. Template Information

**Backend Template (`templates/ddd-api/`):**
- **Example Entity**: "item" - a simple CRUD example
- **Clean Architecture**: Repository, Service, Handler layers
- **Error Handling**: Professional error utilities with `errorutils.AnalyzeDBErr()` - MUST be used in ALL repository methods
- **SQLX Integration**: Uses sqlx for cleaner database operations
- **No Template Variables in Code**: All Go code is ready to run
- **SPECIFICATION.md**: Generated with project requirements and features (WHAT we're building)
- **AGENTS.md**: Test agent persona for QA coverage with `/test` command

**Frontend Template (`templates/nextjs-frontend/`):**
- **Example Entity**: "item" - matches backend example
- **Next.js 15**: App Router with TypeScript
- **TanStack Query**: Data fetching with React Query
- **Tailwind + shadcn/ui**: Modern styling and components
- **Zustand**: State management
- **No Template Variables in Code**: All React code is ready to run
- **SPECIFICATION.md**: Frontend requirements and UI/UX specifications
- **AGENTS.md**: Frontend test agent for component testing

**Template Files Processed:**
- Backend: `docker-compose.yml.tmpl`, `.env.example.tmpl`, `SPECIFICATION.md.tmpl`
- Frontend: `package.json.tmpl`, `.env.example.tmpl`, `CLAUDE.md.tmpl`, `SPECIFICATION.md.tmpl`

## Example User Prompts

**Simple Request:**
"Create a todo app API without authentication"

**With Frontend:**
"Create a task management app with frontend and authentication"

**Detailed Request:**
"Generate an e-commerce product management API with JWT auth, S3 image uploads, and React frontend"

**Your Response Pattern:**
1. Start with greeting from `prompts/PROJECT_GENERATION_PROMPTS.md`
2. Ask for project name, entity, auth, S3, **frontend**, description
3. Validate inputs and confirm configuration
4. Execute generation process for backend and frontend (if requested)
5. Provide success message with next steps

## IMPORTANT: Always Ask About Frontend

**You MUST always ask about frontend** - this is a key feature that users often want but may not explicitly mention. Even if they just say "create an API", ask if they want a frontend too.

**Example:**
User: "Create a todo API"
You: [Ask about frontend] "Do you also want a Next.js frontend for this API?"

## Quick Generation Instructions for Claude

When user requests a project **with frontend**:

1. **Generate backend first** using existing workflow
2. **After backend is complete**, generate frontend:
   ```bash
   # Run this from generator/ directory
   ./scripts/generate-frontend.sh PROJECT_NAME ENTITY API_PORT FRONTEND_PORT AUTH S3
   ```
3. **Provide both setup instructions** to the user

**Example Full Workflow:**
- User wants: "todo app with frontend and auth"
- App container created at: `/Users/kranti/Documents/Code/Go/todo-app/`
- Backend generated at: `/Users/kranti/Documents/Code/Go/todo-app/todo-app-server/`
- Frontend generated at: `/Users/kranti/Documents/Code/Go/todo-app/todo-app-client/`
- Backend ports: API=8050, DB=5450, Redis=6380
- Frontend port: 3050
- Run: `./scripts/generate-frontend.sh todo-app task 8050 3050 true false`

## Configuration

### What's Preset (You Don't Need to Ask)
These are configured in `generator/config.yaml` and apply to ALL projects:

**Fixed Settings:**
- Database credentials: Always `user`/`password`
- Module prefix: `github.com/darkphotonKN/` (creates module like `github.com/darkphotonKN/project-name`)
- Output location: `../` (creates sibling directories)
- Port base values: API 8000, DB 5432, Redis 6379
- Port randomization: ±50 range for each service
- Git initialization: Always enabled
- **Module Name**: Automatically generated as `{module_prefix}{project-name}` and applied to go.mod and all import statements

**Default Features (but still ask to confirm):**
- Authentication: `true` by default
- S3 uploads: `false` by default
- Frontend: `false` by default
- Redis caching: Always included (never ask)

### What Claude Prompts For (Per Project)
You MUST gather these for each project:

1. **Project Name** - Required (can infer from request)
2. **Primary Entity** - Required (must ask if not clear)
3. **Authentication** - Ask yes/no (default: yes)
4. **S3 Uploads** - Ask yes/no (default: no)
5. **Frontend** - Ask yes/no (default: no)
6. **Description** - Required for documentation

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

## SPECIFICATION.md Generation

When generating a new project, Claude should:

### Mode 1: Source Document Provided
If user provides existing documentation (PDF, markdown, requirements doc):
1. Parse the source document
2. Extract requirements, features, domain terms
3. Map to SPECIFICATION.md sections
4. Flag gaps with [TBD - needs clarification]
5. Keep content concise—summarize, don't copy verbatim

### Mode 2: Starting Fresh (No Source)
If generating from scratch via conversation:
1. Ask user about project type, users, entities, integrations
2. Generate minimal SPECIFICATION.md with 3-5 core features
3. Mark non-essential sections as [TBD]
4. Iterate with user to expand as needed

**IMPORTANT**: SPECIFICATION.md should be generated BEFORE writing any code. It defines WHAT we're building, which guides HOW we build it.

## CRITICAL: Error Handling Standards for Generated Projects

**ALL generated projects include `internal/utils/errorutils` package with:**
- `AnalyzeDBErr()` - Converts database errors to standard application errors
- Standard error types: `ErrNotFound`, `ErrDuplicateResource`, `ErrConstraintViolation`, etc.

**BASELINE REQUIREMENT**: When generating new code for any project:
1. **Repository Layer**: ALWAYS use `errorutils.AnalyzeDBErr(err)` for ALL database operations
2. **Handler Layer**: ALWAYS handle errorutils errors with appropriate HTTP status codes
3. **Service Layer**: Use errorutils standard errors for business logic validation

**This is NON-NEGOTIABLE** - The errorutils package exists in every project specifically to avoid repetitive error checking. Use it!

## Project Documentation Structure

**Generated projects include three key documentation files:**

1. **SPECIFICATION.md**: Project requirements, features, and business rules (WHAT we're building)
   - Generated from template with project-specific details
   - Defines domain terms, API surface, data models
   - Should be updated when requirements change

2. **CLAUDE.md**: General coding standards, patterns, and conventions (HOW we build)
   - Implementation guidelines and best practices
   - Error handling patterns, architecture rules
   - Consistent across all projects using this template

3. **AGENTS.md**: Test agent persona for automated QA coverage (TEST)
   - Invoke with `/test` command in Claude
   - Provides QA engineer mindset for writing tests
   - Coverage priorities and testing boundaries

**Documentation Hierarchy:**
```
SPECIFICATION.md (WHAT) ← Requirements and features
    ↓
CLAUDE.md (HOW) ← Code patterns and standards
    ↓
AGENTS.md (TEST) ← QA persona for test coverage
```

When working with generated projects:
- Check SPECIFICATION.md for requirements before implementing
- Follow coding standards in CLAUDE.md
- Use `/test` command to invoke test agent from AGENTS.md
- Update SPECIFICATION.md when requirements change

## Documentation Reference

- `docs/CLAUDE_GENERATION_WORKFLOW.md` - Complete step-by-step workflow
- `prompts/PROJECT_GENERATION_PROMPTS.md` - Standardized user prompts
- `docs/VALIDATION_CHECKLIST.md` - Input validation rules
- `examples/EXAMPLE_CONVERSATIONS.md` - Example interactions
- `docs/TROUBLESHOOTING_GUIDE.md` - Error resolution

## IMPORTANT: Working with Generated Projects

When users ask you to work on a project that was generated using this template generator, you MUST:

1. **Check for docs/ directory first**:
   - Look for `docs/schema/*.md` for database specifications
   - Look for `docs/api/*.md` for API endpoint specs
   - Look for `docs/integrations/*.md` for third-party service specs
   - Look for `docs/plans/*.md` for previously generated implementation plans

2. **Never guess specifications**:
   - If docs exist, use the EXACT field types, constraints, and API shapes specified
   - The docs/schema/ folder contains one file per table with exact field specs
   - If docs don't exist for what you're implementing, ask before proceeding
   - If there are sensible defaults, you can propose them while asking for confirmation

3. **Follow the documentation hierarchy**:
   - `SPECIFICATION.md` - What we're building (requirements)
   - `CLAUDE.md` - How we build (patterns and standards)
   - `docs/*` - Detailed specs for specific features
   - `AGENTS.md` - Test coverage with /test command

This ensures consistency and prevents implementation errors from guessing or assumptions.

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