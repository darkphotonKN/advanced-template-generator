# Claude Generation Troubleshooting Guide

This guide helps Claude diagnose and resolve common issues during project generation.

## Pre-Generation Issues

### Template Not Found

**Problem**: `templates/ddd-api/` directory missing
```
Error: Template directory not found at templates/ddd-api/
```

**Solutions**:
1. Verify you're in the correct generator directory
2. Check if template files were accidentally moved/deleted
3. Re-clone the repository if templates are missing

**Claude Actions**:
```bash
# Check if we're in the right location
pwd
ls -la templates/

# If missing, guide user to correct location
cd /path/to/go-template-generator/generator
```

### Configuration Missing

**Problem**: `generator/config.yaml` not found
```
Error: Configuration file not found
```

**Solutions**:
1. Check if config.yaml exists in generator/ directory
2. Verify file permissions are readable
3. Recreate config if corrupted

### Go Not Installed

**Problem**: `go version` command fails
```
Error: go: command not found
```

**Solutions**:
1. Install Go from https://golang.org/dl/
2. Add Go to PATH environment variable
3. Restart terminal after installation

**Claude Actions**:
```bash
# Check Go installation
go version

# If not found, guide user to install
echo "Please install Go from https://golang.org/dl/ and restart your terminal"
```

## Input Validation Issues

### Invalid Project Name

**Problem**: User provides invalid project name formats

**Common Invalid Formats**:
- `My Project` (spaces, capitals)
- `my_project` (underscores)
- `my-project-` (trailing hyphen)
- `-my-project` (leading hyphen)
- `my--project` (consecutive hyphens)

**Claude Response Pattern**:
```
❌ **Invalid Project Name**

Project name must:
- Be lowercase letters only
- Use hyphens instead of spaces
- Not contain special characters or underscores
- Not start or end with hyphens
- Not have consecutive hyphens

Examples of valid names:
✅ "todo-app"
✅ "inventory-system"
✅ "blog-api"

Please try again with a valid name:
```

### Invalid Entity Name

**Problem**: User provides plural or complex entity names

**Common Invalid Formats**:
- `tasks` (plural - should be "task")
- `shopping-cart` (complex - should be "cart")
- `User` (capitalized)
- `todo-item` (hyphenated)

**Claude Response Pattern**:
```
❌ **Invalid Entity Name**

Entity name must:
- Be singular (not plural)
- Be a simple lowercase noun
- Not contain hyphens or special characters

Examples of valid entities:
✅ "task" (not "tasks")
✅ "product" (not "products")
✅ "user" (not "User")

Please try again with a valid entity name:
```

## Generation Process Issues

### Directory Already Exists

**Problem**: Target directory already exists
```
Error: Directory 'my-project' already exists
```

**Claude Solutions**:
1. **Suggest different name**:
   ```
   The directory 'my-project' already exists.

   Options:
   1. Choose a different name (e.g., 'my-project-v2')
   2. Remove the existing directory: rm -rf my-project
   3. Cancel and check the existing project

   What would you prefer?
   ```

2. **If user wants to remove**:
   ```bash
   # Confirm before removal
   echo "⚠️  This will permanently delete the existing 'my-project' directory."
   echo "Type 'yes' to confirm:"
   # Wait for confirmation, then:
   rm -rf my-project
   ```

### Port Conflicts

**Problem**: Calculated ports are already in use

**Detection**:
```bash
# Check if ports are in use
netstat -ln | grep :8023
lsof -i :8023
```

**Claude Solutions**:
1. **Auto-regenerate ports**: Use the randomization system to try different ports
2. **Manual port selection**: Ask user for preferred port range
3. **Skip port validation**: Continue with potentially conflicting ports

### Permission Issues

**Problem**: Cannot write to target directory
```
Error: Permission denied creating directory
```

**Claude Solutions**:
```bash
# Check current directory permissions
ls -la .

# Suggest alternative location
echo "Try generating in your home directory:"
cd ~ && mkdir projects && cd projects
```

### Git Issues

**Problem**: Git not available or repository issues

**Common Git Errors**:
- `git: command not found`
- `fatal: not a git repository`
- `fatal: unable to access repository`

**Claude Solutions**:
```bash
# Check git availability
git --version

# Initialize without git if unavailable
echo "Git not available, skipping repository initialization"
echo "You can run 'git init' later in your project directory"
```

## Post-Generation Issues

### Go Module Issues

**Problem**: `go mod tidy` fails
```
Error: go: malformed module path
```

**Claude Solutions**:
1. **Fix module path**:
   ```bash
   # Check go.mod content
   cat go.mod

   # Regenerate if corrupted
   rm go.mod
   go mod init github.com/kranti/project-name
   go mod tidy
   ```

### Template Processing Errors

**Problem**: Variables not replaced in generated files
```
Error: Found unreplaced template variables: {{.ProjectName}}
```

**Claude Debugging**:
1. **Check for .tmpl files**:
   ```bash
   find . -name "*.tmpl"
   # Should return empty - all .tmpl extensions should be removed
   ```

2. **Check variable replacement**:
   ```bash
   grep -r "{{.*}}" . --exclude-dir=.git
   # Should not find any template variables
   ```

3. **Regenerate if needed**:
   ```bash
   cd ..
   rm -rf project-name
   # Run generation again
   ```

### Database Connection Issues

**Problem**: Cannot connect to database after setup
```
Error: failed to connect to database
```

**Claude Solutions**:
1. **Check Docker status**:
   ```bash
   docker ps
   # Should show PostgreSQL and Redis containers running
   ```

2. **Verify ports**:
   ```bash
   # Check if database port is accessible
   telnet localhost 5456
   ```

3. **Check .env configuration**:
   ```bash
   cat .env
   # Verify DB_HOST, DB_PORT, DB_USER, DB_PASSWORD match docker-compose.yml
   ```

### Migration Failures

**Problem**: `make migrate-up` fails
```
Error: migration failed
```

**Claude Solutions**:
1. **Check database connectivity**:
   ```bash
   # Test database connection
   docker exec -it {project}_postgres_1 psql -U user -d project_name_db -c "SELECT 1;"
   ```

2. **Check migration files**:
   ```bash
   ls -la migrations/
   # Should contain .up.sql and .down.sql files
   ```

3. **Manual migration**:
   ```bash
   # Run migrations manually
   migrate -path migrations -database "postgres://user:password@localhost:5456/project_name_db?sslmode=disable" up
   ```

## Development Server Issues

### Hot Reload Not Working

**Problem**: Code changes don't trigger restart
```
Error: Air not watching files properly
```

**Claude Solutions**:
1. **Check Air configuration**:
   ```bash
   cat .air.toml
   # Verify include_ext and exclude_dir settings
   ```

2. **Restart Air**:
   ```bash
   # Kill existing Air process
   pkill air
   make dev
   ```

### API Not Responding

**Problem**: Server starts but API doesn't respond
```
Error: connection refused on localhost:8023
```

**Claude Solutions**:
1. **Check server logs**:
   ```bash
   # Look for startup errors in terminal output
   ```

2. **Verify port binding**:
   ```bash
   netstat -ln | grep 8023
   # Should show server listening on port
   ```

3. **Test basic endpoint**:
   ```bash
   curl http://localhost:8023/health
   # Should return basic health check
   ```

## Registry Issues

### Registry File Corruption

**Problem**: `~/.go-gen-projects.json` is corrupted
```
Error: invalid character in registry file
```

**Claude Solutions**:
1. **Backup and reset**:
   ```bash
   mv ~/.go-gen-projects.json ~/.go-gen-projects.json.backup
   echo '{"next_index": 1, "projects": []}' > ~/.go-gen-projects.json
   ```

2. **Validate JSON**:
   ```bash
   cat ~/.go-gen-projects.json | jq .
   # Should parse without errors
   ```

### Port Conflicts in Registry

**Problem**: Registry shows conflicting port assignments

**Claude Solutions**:
1. **Audit registry**:
   ```bash
   cat ~/.go-gen-projects.json | jq '.projects[] | {name, api_port, db_port, redis_port}'
   ```

2. **Regenerate with different ports**:
   ```bash
   # Manual port override if needed
   # Update project configuration
   ```

## Recovery Procedures

### Complete Project Reset

**When to use**: Multiple issues, easier to restart

**Claude Actions**:
```bash
# 1. Remove failed project
rm -rf project-name

# 2. Clean registry entry (if added)
# Edit ~/.go-gen-projects.json to remove project entry

# 3. Regenerate from scratch
# Start generation process again
```

### Partial Recovery

**When to use**: Minor template issues, specific file problems

**Claude Actions**:
```bash
# 1. Identify specific problem files
grep -r "{{.*}}" . --exclude-dir=.git

# 2. Fix individual files
# Manual editing or re-copy from template

# 3. Test functionality
make dev
```

### Environment Reset

**When to use**: Docker or development environment issues

**Claude Actions**:
```bash
# 1. Stop all containers
make docker-down

# 2. Remove containers and volumes
docker-compose down -v

# 3. Restart fresh
make docker-up
make migrate-up
```

## Prevention Best Practices

### For Claude

1. **Always validate inputs** before starting generation
2. **Check prerequisites** (Go, git, docker) before proceeding
3. **Use absolute paths** when possible
4. **Provide specific error messages** with actionable solutions
5. **Test basic functionality** after generation when possible

### For Users

1. **Use valid naming conventions** from the start
2. **Ensure clean working directory** before generation
3. **Have Docker running** before infrastructure setup
4. **Check port availability** if experiencing conflicts
5. **Keep templates unmodified** until familiar with system

## Emergency Contacts

**System Issues**: Check generator repository issues
**Template Problems**: Refer to original cashflow project structure
**Docker Issues**: Verify Docker Desktop is running and updated
**Go Issues**: Verify Go version compatibility (1.19+)

This troubleshooting guide ensures Claude can help users resolve common issues quickly and get their projects running successfully.