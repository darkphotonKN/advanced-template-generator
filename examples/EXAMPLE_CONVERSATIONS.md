# Example Generation Conversations

This document shows complete example conversations between users and Claude for generating different types of projects.

## Example 1: Simple Todo App (No Auth)

### **User Request:**
> "I want to create a simple todo app API without authentication"

### **Claude Response:**
```
🚀 **Go DDD API Project Generator**

I'll help you create a production-ready DDD API project with clean architecture, hot reload, database setup, and best practices.

Let's start by gathering some information about your project:

📝 **Project Name**
What would you like to call your project?

Guidelines:
- Use lowercase letters only
- Hyphens are okay (e.g., "todo-app", "user-service")
- No spaces or special characters

Example: "inventory-system", "blog-api", "task-manager"

Your project name:
```

### **User:** `todo-app`

### **Claude:**
```
🏗️ **Primary Entity**
What's the main "thing" your API will manage?

This becomes your domain model and replaces "financial" from our template.

Examples:
- "task" → /api/tasks endpoints, Task struct, tasks table
- "product" → /api/products endpoints, Product struct, products table

Your primary entity:
```

### **User:** `task`

### **Claude:**
```
🔐 **Authentication**
Do you need JWT authentication for your API?

- **Yes (recommended)**: Adds JWT middleware, login endpoints, protected routes
- **No**: Completely public API, no authentication required

Include authentication? (yes/no, default: yes):
```

### **User:** `no`

### **Claude:**
```
📁 **File Uploads**
Do you need S3 file upload support?

- **Yes**: Adds S3 service, presigned URLs, image upload endpoints
- **No (recommended)**: Simpler setup, no cloud dependencies

Include S3 uploads? (yes/no, default: no):
```

### **User:** `no`

### **Claude:**
```
📋 **Project Description**
Brief description of what this API does:

Example: "Task management API for personal productivity"

Your description:
```

### **User:** `Simple task management for daily todos`

### **Claude:**
```
📋 **Please Confirm Your Configuration:**

**Project Details:**
- Name: todo-app
- Module: github.com/darkphotonKN/todo-app
- Database: todo_app_db

**Entity Configuration:**
- Primary Entity: task
- API Endpoints: /api/tasks
- Database Table: tasks
- Go Struct: Task

**Features:**
- Authentication: ❌
- S3 Uploads: ❌
- Redis Caching: ✅ (always included)

**Infrastructure:**
- API Port: ~8010±50 (randomized)
- Database Port: ~5442±50 (randomized)
- Redis Port: ~6389±50 (randomized)

Does this look correct? Type 'yes' to proceed:
```

### **User:** `yes`

### **Claude:**
```
✅ Generating your todo-app project...

📁 Copying template structure...
🔧 Processing template variables...
🏗️ Renaming entity files...
🐹 Initializing Go module...
🔄 Setting up git repository...
📋 Updating project registry...

🎉 **Project Generated Successfully!**

Your DDD API project "todo-app" has been created with:

📁 **Project Structure:**
```
todo-app/
├── cmd/main.go                    # Application entry point
├── internal/task/                 # Your task domain
│   ├── model.go                   # Task struct & requests
│   ├── handler.go                 # HTTP endpoints
│   ├── service.go                 # Business logic
│   └── repository.go              # Database operations
├── config/routes.go               # API routes (no auth)
├── migrations/                    # Database schema
├── docker-compose.yml            # Infrastructure setup
└── CLAUDE.md                     # AI development guide
```

🔢 **Assigned Ports:**
- API Server: http://localhost:8023
- Database: localhost:5456
- Redis: localhost:6401

🚀 **Next Steps:**
```bash
cd todo-app
cp .env.example .env
make docker-up     # Start PostgreSQL & Redis
make migrate-up    # Create tasks table
make dev          # Start API with hot reload
```

🌐 **Your API Endpoints:**
- `GET /api/tasks` - List all tasks
- `POST /api/tasks` - Create new task
- `GET /api/tasks/{id}` - Get specific task
- `PUT /api/tasks/{id}` - Update task
- `DELETE /api/tasks/{id}` - Delete task

Ready to build something awesome! 🚀
```

---

## Example 2: E-commerce API with Auth and S3

### **User Request:**
> "Create an e-commerce product management API with authentication and image uploads"

### **Claude Response:**
```
🚀 **Go DDD API Project Generator**

I'll help you create a production-ready DDD API project with clean architecture, hot reload, database setup, and best practices.

📝 **Project Name**
Your project name:
```

### **User:** `ecommerce-api`

### **Claude:**
```
🏗️ **Primary Entity**
Your primary entity:
```

### **User:** `product`

### **Claude:**
```
🔐 **Authentication**
Include authentication? (yes/no, default: yes):
```

### **User:** `yes`

### **Claude:**
```
📁 **File Uploads**
Include S3 uploads? (yes/no, default: no):
```

### **User:** `yes`

### **Claude:**
```
📋 **Project Description**
Your description:
```

### **User:** `E-commerce product catalog with image management`

### **Claude:**
```
📋 **Please Confirm Your Configuration:**

**Project Details:**
- Name: ecommerce-api
- Module: github.com/darkphotonKN/ecommerce-api
- Database: ecommerce_api_db

**Entity Configuration:**
- Primary Entity: product
- API Endpoints: /api/products
- Database Table: products
- Go Struct: Product

**Features:**
- Authentication: ✅
- S3 Uploads: ✅
- Redis Caching: ✅

**Infrastructure:**
- API Port: ~8020±50 (randomized)
- Database Port: ~5452±50 (randomized)
- Redis Port: ~6399±50 (randomized)

Does this look correct? Type 'yes' to proceed:
```

### **User:** `yes`

### **Claude:**
```
🎉 **Project Generated Successfully!**

Your DDD API project "ecommerce-api" has been created with:

📁 **Enhanced Features:**
- JWT Authentication middleware
- S3 file upload service
- Image upload endpoints
- Product image management
- Protected API routes

🔢 **Assigned Ports:**
- API Server: http://localhost:8067
- Database: localhost:5489
- Redis: localhost:6384

🔑 **Authentication Endpoints:**
- `POST /api/public/register` - User registration
- `POST /api/public/login` - User login

🌐 **Product Endpoints (Requires Auth):**
- `GET /api/products` - List all products
- `POST /api/products` - Create new product
- `GET /api/products/{id}` - Get specific product
- `PUT /api/products/{id}` - Update product
- `DELETE /api/products/{id}` - Delete product

📸 **Image Upload Endpoints:**
- `POST /api/uploads/request` - Request presigned URL
- `GET /api/uploads/{id}/status` - Check upload status

🚀 **Next Steps:**
```bash
cd ecommerce-api
cp .env.example .env
# Configure your AWS S3 credentials in .env
make docker-up
make migrate-up
make dev
```

💡 **S3 Setup Required:**
Add these to your `.env` file:
- `AWS_REGION=us-east-1`
- `AWS_ACCESS_KEY_ID=your-key`
- `AWS_SECRET_ACCESS_KEY=your-secret`
- `S3_BUCKET_NAME=your-bucket`
```

---

## Example 3: Error Handling - Invalid Names

### **User Request:**
> "Create a Blog App API"

### **Claude Response:**
```
📝 **Project Name**
Your project name:
```

### **User:** `Blog App`

### **Claude:**
```
❌ **Invalid Project Name**

Project name must:
- Be lowercase letters only
- Use hyphens instead of spaces
- Not contain special characters

Examples of valid names:
✅ "blog-app"
✅ "blog-api"
✅ "personal-blog"

Examples of invalid names:
❌ "Blog App" (has space and capitals)
❌ "blog_app" (underscore not allowed)

Please try again with a valid name:
```

### **User:** `blog-api`

### **Claude:**
```
🏗️ **Primary Entity**
Your primary entity:
```

### **User:** `posts`

### **Claude:**
```
❌ **Invalid Entity Name**

Entity name must:
- Be singular (not plural)
- Be a simple lowercase noun

Examples of valid entities:
✅ "post" (not "posts")
✅ "article" (not "articles")

Please try again with a valid entity name:
```

### **User:** `post`

### **Claude:**
```
✅ Perfect! Continuing with entity: post

🔐 **Authentication**
Include authentication? (yes/no, default: yes):
```

---

## Example 4: Quick Generation

### **User Request:**
> "Generate inventory management API with product entity"

### **Claude:**
```
🚀 **Go DDD API Project Generator**

I can see you want an inventory management API with product entity. Let me gather a few more details:

📝 **Project Name**: inventory-api ✅ (inferred)
🏗️ **Primary Entity**: product ✅ (specified)

🔐 **Authentication**
Include authentication? (yes/no, default: yes):
```

### **User:** `yes`

### **Claude:**
```
📁 **File Uploads**
Include S3 uploads? (yes/no, default: no):
```

### **User:** `no`

### **Claude:**
```
📋 **Project Description**
Brief description (or press enter for default):
```

### **User:** *[presses enter]*

### **Claude:**
```
📋 **Please Confirm Your Configuration:**

**Project Details:**
- Name: inventory-api
- Module: github.com/darkphotonKN/inventory-api
- Database: inventory_api_db
- Description: DDD API for product management

**Entity Configuration:**
- Primary Entity: product
- API Endpoints: /api/products

**Features:**
- Authentication: ✅
- S3 Uploads: ❌

Does this look correct? Type 'yes' to proceed:
```

### **User:** `yes`

### **Claude:**
```
🎉 **Project "inventory-api" Generated Successfully!**

🔢 **Assigned Ports:**
- API Server: http://localhost:8041
- Database: localhost:5467
- Redis: localhost:6412

🚀 **Quick Start:**
```bash
cd inventory-api && cp .env.example .env && make docker-up && make migrate-up && make dev
```

Your inventory API is ready at http://localhost:8041! 🚀
```

These examples show how Claude should handle different scenarios, validate inputs, and provide helpful error messages while maintaining a friendly, efficient interaction flow.