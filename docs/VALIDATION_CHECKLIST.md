# Claude Generation Validation Checklist

This checklist ensures Claude validates all inputs and configurations before generating a project.

## Pre-Generation Validation

### ✅ Template Structure Check
- [ ] `templates/ddd-api/` directory exists
- [ ] Template contains all required `.tmpl` files
- [ ] `generator/config.yaml` exists and is readable
- [ ] Base migration files exist in templates

### ✅ System Requirements
- [ ] Go is installed (`go version` works)
- [ ] Git is available (`git --version` works)
- [ ] Current directory is writable

## Input Validation

### ✅ Project Name Validation
```bash
# Valid format check
[a-z][a-z0-9-]*[a-z0-9]
```

**Rules:**
- [ ] Contains only lowercase letters, numbers, and hyphens
- [ ] Starts with a letter
- [ ] Ends with letter or number (not hyphen)
- [ ] No consecutive hyphens
- [ ] Length between 2-50 characters
- [ ] Not a reserved word (go, main, test, etc.)

**Valid Examples:** ✅
- `todo-app`
- `inventory-system`
- `blog-api`
- `user-service`

**Invalid Examples:** ❌
- `Todo-App` (contains uppercase)
- `todo_app` (contains underscore)
- `todo--app` (consecutive hyphens)
- `-todo` (starts with hyphen)
- `app-` (ends with hyphen)

### ✅ Entity Name Validation
```bash
# Valid format check
[a-z][a-z0-9]*
```

**Rules:**
- [ ] Contains only lowercase letters and numbers
- [ ] Starts with a letter
- [ ] Singular form (not plural)
- [ ] Simple noun (no complex phrases)
- [ ] Length between 2-20 characters
- [ ] Not a Go keyword (type, func, var, etc.)

**Valid Examples:** ✅
- `task`
- `product`
- `user`
- `order`
- `post`

**Invalid Examples:** ❌
- `tasks` (plural)
- `shopping-cart` (contains hyphen)
- `User` (contains uppercase)
- `type` (Go keyword)
- `todo-item` (too complex)

### ✅ Directory Validation
- [ ] Target directory doesn't already exist
- [ ] Parent directory is writable
- [ ] Sufficient disk space (at least 50MB)

## Configuration Validation

### ✅ Port Availability Check
```bash
# Check if ports are available (optional but helpful)
netstat -ln | grep :{calculated_port}
```
- [ ] Calculated API port is not in use
- [ ] Calculated DB port is not in use
- [ ] Calculated Redis port is not in use
- [ ] All ports are within valid range (1024-65535)

### ✅ Registry Validation
- [ ] Can read/write to `~/.go-gen-projects.json`
- [ ] Project name doesn't already exist in registry
- [ ] Registry format is valid JSON

## Template Variable Validation

### ✅ Required Variables
- [ ] `{{.ProjectName}}` → Valid project name
- [ ] `{{.ModuleName}}` → github.com/kranti/{project-name}
- [ ] `{{.PrimaryEntity}}` → Valid entity name
- [ ] `{{.EntityCapitalized}}` → Capitalized entity
- [ ] `{{.EntityPlural}}` → Proper plural form
- [ ] `{{.APIPort}}` → Valid port number string
- [ ] `{{.DBPort}}` → Valid port number string
- [ ] `{{.RedisPort}}` → Valid port number string
- [ ] `{{.DBName}}` → {project_name}_db format
- [ ] `{{.DBUser}}` → "user" (from config)
- [ ] `{{.DBPassword}}` → "password" (from config)
- [ ] `{{.IncludeAuth}}` → true/false boolean
- [ ] `{{.IncludeS3}}` → true/false boolean
- [ ] `{{.ProjectDescription}}` → Non-empty string

### ✅ Derived Values Check
- [ ] Database name uses underscores (not hyphens)
- [ ] Module name is valid Go module format
- [ ] Entity plural follows English rules
- [ ] All boolean flags are properly set

## Pre-Execution Validation

### ✅ Template Processing Test
- [ ] All template files can be parsed
- [ ] No undefined variables in templates
- [ ] No syntax errors in template logic
- [ ] Conditional includes work properly

### ✅ File System Check
- [ ] Can create target directory
- [ ] Can copy template files
- [ ] Can write processed files
- [ ] Can execute git commands

## Post-Generation Validation

### ✅ Generated Project Check
- [ ] All template files were processed
- [ ] No `.tmpl` extensions remain
- [ ] Entity directory renamed correctly
- [ ] `go.mod` file created successfully
- [ ] Git repository initialized
- [ ] No syntax errors in generated Go files

### ✅ Configuration Files
- [ ] `docker-compose.yml` has correct ports
- [ ] `.env.example` has correct values
- [ ] `CLAUDE.md` has project description
- [ ] Migration files use correct table names

### ✅ Registry Update
- [ ] Project added to registry file
- [ ] Port assignments recorded
- [ ] Next index incremented
- [ ] Registry file format preserved

## Error Recovery Checklist

### ✅ Cleanup on Failure
- [ ] Remove partially created directory
- [ ] Don't update registry if generation failed
- [ ] Preserve original template files
- [ ] Log specific error for debugging

### ✅ Common Error Solutions
- [ ] **"Directory exists"** → Suggest different name or removal
- [ ] **"Invalid name"** → Show format rules and examples
- [ ] **"Port in use"** → Generate different random ports
- [ ] **"Template missing"** → Verify generator installation
- [ ] **"Go not found"** → Ask user to install Go
- [ ] **"Git not found"** → Skip git init with warning

## Quality Assurance

### ✅ Generated Project Quality
- [ ] Project builds successfully (`go build`)
- [ ] No Go syntax errors
- [ ] No import path errors
- [ ] Docker compose starts correctly
- [ ] Migrations run successfully
- [ ] Basic API endpoints respond

### ✅ Documentation Quality
- [ ] CLAUDE.md contains project-specific info
- [ ] README has correct commands
- [ ] All placeholder text replaced
- [ ] Entity names used consistently throughout

This checklist ensures every generated project meets quality standards and follows all established conventions.