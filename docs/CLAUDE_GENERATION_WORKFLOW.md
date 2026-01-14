# Claude-Assisted Project Generation Workflow

This document provides step-by-step instructions for Claude to help users generate new DDD API projects using the go-template-generator.

## Pre-Generation Setup

Before starting any generation, Claude should:
1. Verify the template structure exists at `templates/ddd-api/`
2. Check that `generator/config.yaml` contains the permanent settings
3. Load the current project registry from `~/.go-gen-projects.json` if it exists

## Interactive Generation Workflow

### Phase 1: Requirements Gathering

**Claude should ask these questions in order:**

```
🚀 Let's create a new DDD API project!

1. **Project Name**: What would you like to call your project?
   - Must be lowercase, can include hyphens
   - Will become directory name and part of module path
   - Example: "todo-app", "inventory-system", "blog-api"

2. **Primary Entity**: What's the main thing your API will manage?
   - This becomes your domain model (replaces "financial" from template)
   - Should be singular, lowercase
   - Example: "task", "product", "post", "user"

3. **Authentication**: Do you need JWT authentication?
   - Default: Yes (recommended for most APIs)
   - Say "no" only if it's a completely public API

4. **File Uploads**: Do you need S3 file upload support?
   - Default: No (adds complexity)
   - Say "yes" if you need image/document uploads

5. **Description**: Brief description of your project
   - Will be added to CLAUDE.md
   - Example: "Task management API for personal productivity"
```

### Phase 2: Validation & Confirmation

**Claude should validate and confirm:**

```
📋 Project Configuration:
- Name: {project-name}
- Entity: {entity} (plural: {entities})
- Module: github.com/kranti/{project-name}
- Database: {project_name}_db
- Authentication: {yes/no}
- S3 Uploads: {yes/no}
- Assigned Ports: API ~{port}±50, DB ~{port}±50, Redis ~{port}±50

Does this look correct? (yes/no)
```

### Phase 3: Generation Execution

**Claude should execute these steps:**

1. **Copy Template Structure**
   ```bash
   cp -r templates/ddd-api/ {project-name}/
   cd {project-name}/
   ```

2. **Calculate Ports** (using randomization logic from config.yaml)
   - Load next project index from registry
   - Calculate base ports: (index × 10) + base_port
   - Apply random offset ±50 for each service
   - Ensure ports are within 1024-65535 range

3. **Process Template Variables**
   Replace these variables in all `.tmpl` files:
   ```
   {{.ProjectName}} → project-name
   {{.ModuleName}} → github.com/kranti/project-name
   {{.PrimaryEntity}} → entity
   {{.EntityCapitalized}} → Entity
   {{.EntityPlural}} → entities
   {{.APIPort}} → calculated_api_port
   {{.DBPort}} → calculated_db_port
   {{.RedisPort}} → calculated_redis_port
   {{.DBName}} → project_name_db
   {{.DBUser}} → user
   {{.DBPassword}} → password
   {{.IncludeAuth}} → true/false
   {{.IncludeS3}} → true/false
   {{.ProjectDescription}} → user_provided_description
   ```

4. **Rename and Clean Files**
   - Remove `.tmpl` extensions from all template files
   - Rename `internal/entity/` to `internal/{entity}/`
   - Process conditional includes (remove auth middleware if --no-auth)

5. **Initialize Project**
   ```bash
   go mod init github.com/kranti/{project-name}
   go mod tidy
   rm -rf .git
   git init
   git add .
   git commit -m "initial commit"
   ```

6. **Update Registry**
   Add project to `~/.go-gen-projects.json`:
   ```json
   {
     "name": "project-name",
     "index": next_index,
     "api_port": calculated_port,
     "db_port": calculated_port,
     "redis_port": calculated_port,
     "entity": "entity",
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
└── CLAUDE.md                     # Development guide

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

💡 Tip: Check CLAUDE.md in your project for development patterns and best practices!
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