package ddd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kranti/go-template-generator/internal/config"
	"github.com/kranti/go-template-generator/internal/git"
	"github.com/kranti/go-template-generator/internal/ports"
	"github.com/kranti/go-template-generator/internal/registry"
)

type GeneratorOptions struct {
	ProjectName        string
	Entity             string
	IncludeAuth        bool
	IncludeS3          bool
	IncludeRedis       bool
	ProjectDescription string
	Config             *config.Config
	APIPort            int
	DBPort             int
	RedisPort          int
}

type Generator struct {
	opts      *GeneratorOptions
	registry  *registry.Manager
	portMgr   *ports.Manager
	templateDir string
	targetDir   string
}

func NewGenerator(opts *GeneratorOptions) *Generator {
	// Calculate template directory (relative to current working directory)
	cwd, _ := os.Getwd()
	var templateDir string

	// Check if we're in the generator directory
	if strings.HasSuffix(cwd, "/generator") {
		templateDir = filepath.Join(cwd, "..", "templates", "ddd-api")
	} else if strings.HasSuffix(cwd, "/go-template-generator") {
		// We're in the parent directory
		templateDir = filepath.Join(cwd, "templates", "ddd-api")
	} else {
		// Fallback to relative to binary
		execPath, _ := os.Executable()
		generatorDir := filepath.Dir(execPath)
		templateDir = filepath.Join(generatorDir, "..", "templates", "ddd-api")
	}

	return &Generator{
		opts:        opts,
		registry:    registry.NewManager(opts.Config.ProjectsRegistry),
		portMgr:     ports.NewManager(opts.Config),
		templateDir: templateDir,
		targetDir:   opts.ProjectName,
	}
}

func (g *Generator) Generate() error {
	// Check if project already exists
	exists, err := g.registry.ProjectExists(g.opts.ProjectName)
	if err != nil {
		return fmt.Errorf("failed to check if project exists: %w", err)
	}
	if exists {
		return fmt.Errorf("project '%s' already exists", g.opts.ProjectName)
	}

	// Check if directory already exists
	if _, err := os.Stat(g.targetDir); !os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' already exists", g.targetDir)
	}

	// Get next project index and allocate ports
	nextIndex, err := g.registry.GetNextIndex()
	if err != nil {
		return fmt.Errorf("failed to get next project index: %w", err)
	}

	allocatedPorts := g.portMgr.AllocatePorts(nextIndex)
	g.opts.APIPort = allocatedPorts.API
	g.opts.DBPort = allocatedPorts.DB
	g.opts.RedisPort = allocatedPorts.Redis

	// Copy template
	fmt.Printf("📁 Creating project directory '%s'...\n", g.opts.ProjectName)
	if err := g.copyTemplate(); err != nil {
		return fmt.Errorf("failed to copy template: %w", err)
	}

	// Generate template variables
	vars := g.generateTemplateVars()

	// Process templates
	fmt.Printf("🔧 Processing templates...\n")
	replacer := NewReplacer(vars)
	if err := replacer.ProcessDirectory(g.targetDir); err != nil {
		return fmt.Errorf("failed to process templates: %w", err)
	}

	// Rename template files
	fmt.Printf("📝 Finalizing files...\n")
	if err := replacer.RenameTemplateFiles(g.targetDir); err != nil {
		return fmt.Errorf("failed to rename template files: %w", err)
	}

	// Rename entity directory
	if err := replacer.RenameEntityDirectory(g.targetDir); err != nil {
		return fmt.Errorf("failed to rename entity directory: %w", err)
	}

	// Clean up Go imports
	if err := replacer.CleanGoImports(g.targetDir); err != nil {
		return fmt.Errorf("failed to clean Go imports: %w", err)
	}

	// Initialize Go module
	fmt.Printf("🐹 Initializing Go module...\n")
	if err := g.initGoModule(vars.ModuleName); err != nil {
		return fmt.Errorf("failed to initialize Go module: %w", err)
	}

	// Initialize git repository
	fmt.Printf("🔄 Initializing git repository...\n")
	gitMgr := git.NewManager(g.targetDir)
	if gitMgr.IsGitAvailable() {
		if err := gitMgr.Initialize(g.opts.Config.Git.InitialCommitMessage); err != nil {
			fmt.Printf("⚠️  Warning: failed to initialize git repository: %v\n", err)
		}
	} else {
		fmt.Printf("⚠️  Warning: git not found, skipping git initialization\n")
	}

	// Register project
	fmt.Printf("📋 Registering project...\n")
	if err := g.registry.AddProject(g.opts.ProjectName, g.opts.Entity,
		g.opts.APIPort, g.opts.DBPort, g.opts.RedisPort); err != nil {
		return fmt.Errorf("failed to register project: %w", err)
	}

	return nil
}

func (g *Generator) copyTemplate() error {
	return filepath.Walk(g.templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate relative path
		relPath, err := filepath.Rel(g.templateDir, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(g.targetDir, relPath)

		if info.IsDir() {
			// Create directory
			return os.MkdirAll(targetPath, info.Mode())
		}

		// Copy file
		return g.copyFile(path, targetPath, info.Mode())
	})
}

func (g *Generator) copyFile(src, dst string, mode os.FileMode) error {
	// Create destination directory if it doesn't exist
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	// Read source file
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// Write to destination
	return os.WriteFile(dst, data, mode)
}

func (g *Generator) generateTemplateVars() *TemplateVars {
	entityCapitalized, entityPlural := GenerateEntityNames(g.opts.Entity)

	// Generate module name
	moduleName := g.opts.Config.Defaults.ModulePrefix + g.opts.ProjectName

	// Generate database name (replace hyphens with underscores)
	dbName := strings.ReplaceAll(g.opts.ProjectName, "-", "_") + "_db"

	return &TemplateVars{
		ProjectName:        g.opts.ProjectName,
		ModuleName:         moduleName,
		PrimaryEntity:      g.opts.Entity,
		EntityCapitalized:  entityCapitalized,
		EntityPlural:       entityPlural,
		APIPort:            fmt.Sprintf("%d", g.opts.APIPort),
		DBPort:             fmt.Sprintf("%d", g.opts.DBPort),
		RedisPort:          fmt.Sprintf("%d", g.opts.RedisPort),
		DBName:             dbName,
		DBUser:             g.opts.Config.Database.User,
		DBPassword:         g.opts.Config.Database.Password,
		IncludeAuth:        g.opts.IncludeAuth,
		IncludeS3:          g.opts.IncludeS3,
		IncludeRedis:       g.opts.IncludeRedis,
		ProjectDescription: g.opts.ProjectDescription,
	}
}

func (g *Generator) initGoModule(moduleName string) error {
	originalDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.Chdir(g.targetDir); err != nil {
		return err
	}
	defer os.Chdir(originalDir)

	// Initialize Go module
	cmd := exec.Command("go", "mod", "init", moduleName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run 'go mod init': %w", err)
	}

	// Run go mod tidy
	cmd = exec.Command("go", "mod", "tidy")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run 'go mod tidy': %w", err)
	}

	return nil
}

// NewProjectRegistry creates a new project registry (used by main.go)
func NewProjectRegistry(registryPath string) *registry.Manager {
	return registry.NewManager(registryPath)
}