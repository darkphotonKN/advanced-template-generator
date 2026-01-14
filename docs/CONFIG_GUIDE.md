# Configuration Guide

This guide explains all configuration options for the go-template-generator system.

## Configuration File Location

Main configuration: `generator/config.yaml`

This file contains permanent settings that apply to ALL generated projects.

## Configuration Sections

### 1. Database Configuration

```yaml
database:
  user: "user"                   # Database username for all projects
  password: "password"           # Database password for all projects
  name_pattern: "{project}_db"  # How database names are generated
```

**When to change:**
- If you prefer different default credentials
- If you want a different database naming pattern

**Impact:**
- Affects ALL generated projects
- Update your local PostgreSQL to match these credentials

### 2. Port Allocation

```yaml
ports:
  base_api: 8000      # Starting port for API servers
  base_db: 5432       # Starting port for databases
  base_redis: 6379    # Starting port for Redis
  increment: 10       # Port increment between projects
  randomization:
    enabled: true     # Enable port randomization
    range: 50         # Random offset range (±50)
```

**How it works:**
1. Project 1: Base ports + (0 × increment) ± random
2. Project 2: Base ports + (1 × increment) ± random
3. Project 3: Base ports + (2 × increment) ± random

**Example:**
- Project 1 API: 8000 + random(-50 to 50) = 7950-8050
- Project 2 API: 8010 + random(-50 to 50) = 7960-8060

**When to change:**
- If default ports conflict with your services
- To disable randomization (set `enabled: false`)
- To adjust random range for more/less variation

### 3. Project Defaults

```yaml
defaults:
  module_prefix: "github.com/kranti/"  # Go module prefix
  include_auth: true                   # Default auth setting
  include_redis: true                  # Redis always included
  include_s3: false                    # Default S3 setting
  primary_entity: "item"               # Fallback entity name
```

**When to change:**
- **module_prefix**: Change to your GitHub username
- **include_auth**: If most projects don't need auth
- **include_s3**: If most projects need file uploads
- **primary_entity**: Different default entity name

### 4. Output Location

```yaml
output_location: "../"  # Where projects are created
```

**Current behavior:**
```
/your-go-projects/
├── go-template-generator/    # This repository
├── todo-app/                # Generated here (../)
└── blog-api/               # Generated here (../)
```

**Alternative configurations:**
- `"."` - Create in current directory
- `"/Users/kranti/projects/"` - Absolute path
- `"../../generated/"` - Different relative path

### 5. Project Registry

```yaml
projects_registry: "~/.go-gen-projects.json"
```

**What it tracks:**
```json
{
  "next_index": 5,
  "projects": [
    {
      "name": "todo-app",
      "index": 1,
      "api_port": 8023,
      "db_port": 5467,
      "redis_port": 6401,
      "entity": "task",
      "created_at": "2024-01-14T10:00:00Z"
    }
  ]
}
```

**When to change:**
- To use different location for registry
- To reset project indexing (delete the file)

### 6. Git Settings

```yaml
git:
  initial_commit_message: "initial commit"
```

**When to change:**
- To use different commit message convention
- To match your team's standards

### 7. Feature Flags

```yaml
features:
  auth:
    enabled: true    # Default value
    description: "JWT authentication middleware"
  s3:
    enabled: false   # Default value
    description: "S3 file upload support"
  redis:
    enabled: true    # Always included
    description: "Redis caching"
```

**Important:**
- These are DEFAULTS only
- Claude will still ask users to confirm
- Redis is always included (DDD template requirement)

## How to Customize for Your Environment

### For Different GitHub Username

Edit `generator/config.yaml`:
```yaml
defaults:
  module_prefix: "github.com/yourusername/"
```

### For Different Database Setup

Edit `generator/config.yaml`:
```yaml
database:
  user: "postgres"
  password: "mysecretpassword"
```

### For Different Port Ranges

Edit `generator/config.yaml`:
```yaml
ports:
  base_api: 3000    # Start from 3000 instead
  base_db: 5500     # Different PostgreSQL range
  base_redis: 6400  # Different Redis range
```

### For Corporate Proxy/Firewall

Edit `generator/config.yaml`:
```yaml
ports:
  randomization:
    enabled: false  # Use predictable ports
```

## What Claude Will Always Ask

Regardless of configuration, Claude will prompt for:

1. **Project Name** - Unique per project
2. **Primary Entity** - Core domain model
3. **Authentication** - Yes/No (shows default)
4. **S3 Uploads** - Yes/No (shows default)
5. **Description** - For documentation

## Precedence Order

1. User's explicit answer to Claude's prompt
2. Default value in `config.yaml`
3. Hardcoded fallback (if any)

## Troubleshooting Configuration

### Port Conflicts

**Problem:** Generated ports conflict with existing services

**Solution:**
```yaml
# Either change base ports
ports:
  base_api: 9000

# Or increase randomization range
ports:
  randomization:
    range: 100  # ±100 instead of ±50
```

### Module Path Issues

**Problem:** Wrong GitHub username in generated projects

**Solution:**
```yaml
defaults:
  module_prefix: "github.com/correctusername/"
```

### Registry Corruption

**Problem:** Registry file is corrupted or has conflicts

**Solution:**
```bash
# Backup and reset
mv ~/.go-gen-projects.json ~/.go-gen-projects.backup.json
# Generator will create new registry on next use
```

## Best Practices

1. **Set module_prefix once** to your GitHub username
2. **Keep database credentials simple** for development
3. **Use port randomization** to avoid conflicts
4. **Don't modify output_location** unless necessary
5. **Let Claude handle per-project settings** via prompts

## Environment-Specific Configurations

### Local Development
Current defaults are optimized for local development.

### CI/CD Pipeline
```yaml
ports:
  randomization:
    enabled: false  # Predictable ports for testing
```

### Team Sharing
```yaml
defaults:
  module_prefix: "github.com/company/"
database:
  user: "dev_user"
  password: "dev_password"
```

This configuration system balances flexibility with simplicity, allowing customization where needed while maintaining consistent defaults.