# Claude-Assisted Project Generation Workflow

This document provides step-by-step instructions for Claude to help users generate new DDD API projects with optional Next.js frontend using the go-template-generator.

## Pre-Generation Setup

Before starting any generation, Claude should:
1. **Start in root directory** `/go-template-generator/` where CLAUDE.md is located
2. **Navigate to generator**: `cd generator`
3. Verify the template structures exist:
   - Backend: `templates/ddd-api/`
   - Frontend: `templates/nextjs-frontend/`
4. Check that `config.yaml` contains the permanent settings including `output_location: "../"`
5. Load the current project registry from `~/.go-gen-projects.json` if it exists

## Interactive Generation Workflow

### Phase 1: Requirements Gathering

**Claude should ask these questions in order:**

```
🚀 Let's create a new full-stack project!

1. **Project Name**: What would you like to call your project?
   - Must be lowercase, can include hyphens
   - Will become directory name and part of module path
   - Example: "todo-app", "inventory-system", "blog-api"

2. **Primary Entity**: What's the main thing your API will manage?
   - This becomes your domain model (replaces "item" from template)
   - Should be singular, lowercase
   - Example: "task", "product", "post", "user"

3. **Authentication**: Do you need JWT authentication?
   - Default: Yes (recommended for most APIs)
   - Say "no" only if it's a completely public API

4. **File Uploads**: Do you need S3 file upload support?
   - Default: No (adds complexity)
   - Say "yes" if you need image/document uploads

5. **Frontend Application**: Do you want a Next.js frontend?
   - Default: No (API-only project)
   - Say "yes" if you want a complete React UI with forms, tables, etc.
   - Frontend connects automatically to your API

6. **Description**: Brief description of your project
   - Will be added to CLAUDE.md
   - Example: "Task management API for personal productivity"
```

### Phase 2: Validation & Confirmation

**Claude should validate and confirm:**

```
📋 Project Configuration:
- Name: {project-name}
- Entity: {entity} (plural: {entities})
- Module: github.com/darkphotonKN/{project-name}
- Database: {project_name}_db
- Authentication: {yes/no}
- S3 Uploads: {yes/no}
- Frontend: {yes/no}
- Assigned Ports: API ~{port}±50, DB ~{port}±50, Redis ~{port}±50, Frontend ~{port}±50

Generated Projects:
- Backend Only: ../{project-name}/ (Go API)
- Full-Stack: ../{project-name}/ (Container folder)
  - ../{project-name}/{project-name}-server/ (Go API)
  - ../{project-name}/{project-name}-client/ (Next.js App)

Does this look correct? (yes/no)
```

### Phase 3: Generation Execution

**Claude should execute these steps:**

### Backend Generation (Always Done)

1. **Copy Backend Template Structure**
   ```bash
   # From generator/ directory, copy to sibling location
   cp -r templates/ddd-api/ ../{project-name}/
   cd ../{project-name}/
   ```

2. **Calculate Ports** (using randomization logic from config.yaml)
   - Load next project index from registry
   - Calculate base ports: (index × 10) + base_port
   - Apply random offset ±50 for each service
   - Ensure ports are within 1024-65535 range

3. **Process Configuration Template Files**
   Replace variables in these `.tmpl` files:

   **docker-compose.yml.tmpl** → **docker-compose.yml**
   **`.env.example.tmpl`** → **`.env.example`**
   **`SPECIFICATION.md.tmpl`** → **`SPECIFICATION.md`**

   Variables to replace:
   ```
   {{.ProjectName}} → project-name
   {{.ModuleName}} → github.com/darkphotonKN/project-name
   {{.APIPort}} → calculated_api_port
   {{.DBPort}} → calculated_db_port
   {{.RedisPort}} → calculated_redis_port
   {{.DBName}} → project_name_db
   {{.DBUser}} → user
   {{.DBPassword}} → password
   ```

4. **Update Go Module and Dependencies**
   ```bash
   # Update module name in go.mod
   # Update all import paths from "github.com/darkphotonKN/go-template-generator"
   # to "github.com/darkphotonKN/{project-name}"

   # Install dependencies
   go mod tidy
   ```

5. **Initialize Project Repository**
   ```bash
   rm -rf .git
   git init
   git add .
   git commit -m "initial commit"
   ```

6. **Rename Entity References**
   ```bash
   # Replace "item" with user's chosen entity throughout codebase
   # Update file/folder names: internal/item/ → internal/{entity}/
   # Update import paths and struct references
   ```

7. **Documentation Files Created**
   - **SPECIFICATION.md**: Project requirements and features (WHAT)
   - **CLAUDE.md**: Coding standards and patterns (HOW)
   - **AGENTS.md**: Test agent persona for QA (TEST)

### Frontend Generation (If Requested)

8. **Copy Frontend Template Structure**
   ```bash
   # From generator/ directory, copy to same app container folder
   cp -r templates/nextjs-frontend/ ../{project-name}/{project-name}-client/
   cd ../{project-name}/{project-name}-client/
   ```

9. **Process Frontend Template Files**
   Replace variables in these `.tmpl` files:

   **package.json.tmpl** → **package.json**
   **`.env.example.tmpl`** → **`.env.example`**
   **`CLAUDE.md.tmpl`** → **`CLAUDE.md`**
   **`SPECIFICATION.md.tmpl`** → **`SPECIFICATION.md`**

   Variables to replace:
   ```
   {{.ProjectName}} → project-name
   {{.ProjectTitle}} → Project Name (capitalized)
   {{.ProjectDescription}} → user description
   {{.APIPort}} → calculated_api_port
   {{.FrontendPort}} → calculated_frontend_port
   {{.IncludeAuth}} → true/false
   {{.IncludeS3}} → true/false
   ```

10. **Rename Frontend Entity References**
   ```bash
   # Replace "item" with user's chosen entity throughout TypeScript code
   # Update folder names: src/features/item/ → src/features/{entity}/
   # Update file names: use-item.ts → use-{entity}.ts
   # Update component names and imports
   # Update API endpoint references
   ```

11. **Initialize Frontend Dependencies**
    ```bash
    npm install
    git init
    git add .
    git commit -m "initial commit"
    ```

12. **Update Registry**
    Add project to `~/.go-gen-projects.json`:
    ```json
    {
      "name": "project-name",
      "index": next_index,
      "api_port": calculated_port,
      "db_port": calculated_port,
      "redis_port": calculated_port,
      "frontend_port": calculated_port,
      "entity": "entity",
      "has_frontend": true/false,
      "created_at": "2024-01-14T10:00:00Z"
    }
    ```

### Phase 4: Success Confirmation

**Claude should provide:**

```
✅ Project '{project-name}' created successfully!

📁 Project Structure:
├── cmd/main.go                    # Application entry point
├── internal/{entity}/             # Your {entity} domain
├── config/routes.go               # API routes for /{entities}
├── migrations/                    # Database schema
├── docker-compose.yml            # Infrastructure (ports: {db_port}, {redis_port})
├── CLAUDE.md                     # Development guide (HOW)
├── SPECIFICATION.md              # Project requirements (WHAT)
└── AGENTS.md                     # Test agent persona (TEST)

🚀 Next Steps:
1. cd {project-name}
2. cp .env.example .env
3. make docker-up                  # Start PostgreSQL & Redis
4. make migrate-up                 # Create {entities} table
5. make dev                        # Start API server

🌐 Your API will be available at: http://localhost:{api_port}

📡 Available Endpoints:
- GET    /api/{entities}           # List {entities}
- POST   /api/{entities}           # Create {entity}
- GET    /api/{entities}/{id}      # Get specific {entity}
- PUT    /api/{entities}/{id}      # Update {entity}
- DELETE /api/{entities}/{id}      # Delete {entity}

💡 Tips:
- Check SPECIFICATION.md for project requirements and features
- Check CLAUDE.md for development patterns and best practices
- Use `/test` command to invoke test agent from AGENTS.md
```

## Error Handling

**If generation fails, Claude should:**

1. **Check common issues:**
   - Directory already exists
   - Invalid project name format
   - Missing template files
   - Go not installed

2. **Provide specific solutions:**
   - Remove existing directory: `rm -rf {project-name}`
   - Fix naming issues: "Use lowercase with hyphens only"
   - Verify template: "Check templates/ddd-api/ exists"

3. **Offer to retry:**
   - "Would you like to try again with a different name?"
   - "Should I help you fix the issue and regenerate?"

## Customization Workflow

**After successful generation, Claude can help with:**

1. **Model Customization**
   - Add fields to the entity model
   - Update migrations accordingly
   - Modify validation rules

2. **API Extensions**
   - Add custom endpoints
   - Implement business logic
   - Add relationships between entities

3. **Infrastructure Setup**
   - Configure environment variables
   - Set up database connections
   - Customize Docker setup

This workflow ensures consistent, interactive project generation with proper validation and clear next steps!