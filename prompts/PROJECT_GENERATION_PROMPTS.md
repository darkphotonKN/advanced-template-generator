# Claude Project Generation Prompts

This file contains standardized prompts and questions for Claude to use during project generation.

## Initial Project Prompt

```
🚀 **Go DDD API Project Generator**

I'll help you create a production-ready DDD API project with clean architecture, hot reload, database setup, and best practices.

Let's start by gathering some information about your project:
```

## Requirements Gathering Prompts

### 1. Project Name
```
📝 **Project Name**
What would you like to call your project?

Guidelines:
- Use lowercase letters only
- Hyphens are okay (e.g., "todo-app", "user-service")
- No spaces or special characters
- Will become your directory name and module path

Example: "inventory-system", "blog-api", "task-manager"
Note: Project will be created at ../{project-name}/ relative to generator directory

Your project name:
```

### 2. Primary Entity
```
🏗️ **Primary Entity**
What's the main "thing" your API will manage?

This becomes your domain model and replaces "financial" from our template.

Guidelines:
- Use singular form (e.g., "task" not "tasks")
- Lowercase, simple noun
- This will become your database table, API endpoints, and Go structs

Examples:
- "task" → /api/tasks endpoints, Task struct, tasks table
- "product" → /api/products endpoints, Product struct, products table
- "post" → /api/posts endpoints, Post struct, posts table

Your primary entity:
```

### 3. Authentication
```
🔐 **Authentication**
Do you need JWT authentication for your API?

- **Yes (recommended)**: Adds JWT middleware, login endpoints, protected routes
- **No**: Completely public API, no authentication required

Most APIs need auth unless they're completely public (like a read-only blog API).

Include authentication? (yes/no, default: yes):
```

### 4. File Uploads
```
📁 **File Uploads**
Do you need S3 file upload support?

- **Yes**: Adds S3 service, presigned URLs, image upload endpoints
- **No (recommended)**: Simpler setup, no cloud dependencies

Only say yes if you specifically need file/image uploads.

Include S3 uploads? (yes/no, default: no):
```

### 5. Frontend Application
```
🖥️ **Frontend Application**
Do you need a Next.js frontend for your API?

- **Yes**: Adds complete React frontend with forms, tables, auth UI
- **No (recommended)**: API-only project

The frontend will be configured to connect to your API and includes:
- Next.js 15 with TypeScript
- Tailwind CSS and shadcn/ui components
- TanStack Query for data fetching
- Zustand for state management
- Full CRUD UI for your entities

Include frontend? (yes/no, default: no):
```

### 6. Project Description
```
📋 **Project Description**
Brief description of what this API does (will be added to documentation):

Example: "Task management API for personal productivity"
Example: "E-commerce product catalog with inventory tracking"
Example: "Blog platform with posts and comments"

Your description:
```

## Validation Prompts

### Configuration Confirmation
```
📋 **Please Confirm Your Configuration:**

**Project Details:**
- Name: {project-name}
- Module: github.com/kranti/{project-name}
- Database: {project_name}_db

**Entity Configuration:**
- Primary Entity: {entity}
- API Endpoints: /api/{entities}
- Database Table: {entities}
- Go Struct: {EntityCapitalized}

**Features:**
- Authentication: {✅/❌}
- S3 Uploads: {✅/❌}
- Frontend: {✅/❌}
- Redis Caching: ✅ (always included)

**Infrastructure:**
- API Port: ~{calculated_port}±50 (randomized)
- Database Port: ~{calculated_port}±50 (randomized)
- Redis Port: ~{calculated_port}±50 (randomized)
- Frontend Port: ~{calculated_port}±50 (randomized) [if frontend enabled]

Does this look correct? Type 'yes' to proceed or 'no' to start over:
```

### Name Validation Error
```
❌ **Invalid Project Name**

Project name must:
- Be lowercase letters only
- Use hyphens instead of spaces
- Not contain special characters

Examples of valid names:
✅ "todo-app"
✅ "inventory-system"
✅ "blog-api"

Examples of invalid names:
❌ "Todo App" (has space and capitals)
❌ "todo_app" (underscore not allowed)
❌ "todo.app" (dot not allowed)

Please try again with a valid name:
```

### Entity Validation Error
```
❌ **Invalid Entity Name**

Entity name must:
- Be singular (not plural)
- Be a simple lowercase noun
- Not contain spaces or special characters

Examples of valid entities:
✅ "task" (not "tasks")
✅ "product" (not "products")
✅ "user" (not "users")

Examples of invalid entities:
❌ "tasks" (should be "task")
❌ "shopping-cart" (too complex, use "cart")
❌ "User" (should be lowercase)

Please try again with a valid entity name:
```

## Success Prompts

### Generation Complete
```
🎉 **Project Generated Successfully!**

Your DDD API project "{project-name}" has been created with:

📁 **Project Structure:**
```
{project-name}/
├── cmd/main.go                    # Application entry point
├── internal/{entity}/             # Your {entity} domain
│   ├── model.go                   # {EntityCapitalized} struct & requests
│   ├── handler.go                 # HTTP endpoints
│   ├── service.go                 # Business logic
│   └── repository.go              # Database operations
├── config/routes.go               # API routes
├── migrations/                    # Database schema
├── docker-compose.yml            # Infrastructure setup
├── Makefile                      # Development commands
└── CLAUDE.md                     # AI development guide
```

🔢 **Assigned Ports:**
- API Server: http://localhost:{api_port}
- Database: localhost:{db_port}
- Redis: localhost:{redis_port}

🚀 **Next Steps:**
```bash
cd {project-name}
cp .env.example .env
make docker-up     # Start PostgreSQL & Redis
make migrate-up    # Create {entities} table
make dev          # Start API with hot reload
```

🌐 **Your API Endpoints:**
- `GET /api/{entities}` - List all {entities}
- `POST /api/{entities}` - Create new {entity}
- `GET /api/{entities}/{id}` - Get specific {entity}
- `PUT /api/{entities}/{id}` - Update {entity}
- `DELETE /api/{entities}/{id}` - Delete {entity}

💡 **Pro Tips:**
- Check `CLAUDE.md` in your project for development patterns
- Use `make test` to run tests
- Use `make lint` before committing
- API includes hot reload - code changes restart automatically

Ready to build something awesome! 🚀
```

## Follow-up Prompts

### Offer Customization
```
🔧 **Want to customize your project further?**

I can help you:
1. **Add fields** to your {entity} model (e.g., status, priority, due_date)
2. **Create relationships** between entities
3. **Add custom endpoints** beyond basic CRUD
4. **Configure environment** variables
5. **Set up additional** middleware or features

What would you like to customize, or are you ready to start developing?
```

### Troubleshooting Offer
```
🔍 **Need help with setup?**

If you run into any issues:
1. **Docker problems**: Make sure Docker is running
2. **Port conflicts**: Check if ports {api_port}, {db_port}, {redis_port} are available
3. **Go module issues**: Run `go mod tidy` in your project
4. **Database connection**: Check your .env file matches docker-compose.yml

I can help debug any specific errors you encounter!
```

These prompts ensure consistent, friendly interaction while gathering all necessary information for project generation.