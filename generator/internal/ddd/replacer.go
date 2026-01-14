package ddd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

type TemplateVars struct {
	ProjectName        string
	ModuleName         string
	PrimaryEntity      string
	EntityCapitalized  string
	EntityPlural       string
	APIPort            string
	DBPort             string
	RedisPort          string
	DBName             string
	DBUser             string
	DBPassword         string
	IncludeAuth        bool
	IncludeS3          bool
	IncludeRedis       bool
	ProjectDescription string
}

type Replacer struct {
	vars *TemplateVars
}

func NewReplacer(vars *TemplateVars) *Replacer {
	return &Replacer{vars: vars}
}

// ProcessFile processes a single file with template variables
func (r *Replacer) ProcessFile(filePath string) error {
	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Parse and execute template
	tmpl, err := template.New("file").Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse template in %s: %w", filePath, err)
	}

	var result strings.Builder
	if err := tmpl.Execute(&result, r.vars); err != nil {
		return fmt.Errorf("failed to execute template in %s: %w", filePath, err)
	}

	// Write back to file
	if err := os.WriteFile(filePath, []byte(result.String()), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

// ProcessDirectory recursively processes all files in a directory
func (r *Replacer) ProcessDirectory(dirPath string) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-template files
		if info.IsDir() {
			return nil
		}

		// Only process template files (.tmpl) and certain other files
		if r.shouldProcessFile(path) {
			return r.ProcessFile(path)
		}

		return nil
	})
}

// shouldProcessFile determines if a file should be processed
func (r *Replacer) shouldProcessFile(path string) bool {
	// Process .tmpl files and certain other files
	if strings.HasSuffix(path, ".tmpl") {
		return true
	}

	// Also process specific files without .tmpl extension
	fileName := filepath.Base(path)
	processableFiles := []string{
		"docker-compose.yml",
		".env.example",
		"CLAUDE.md",
	}

	for _, file := range processableFiles {
		if fileName == file {
			return true
		}
	}

	return false
}

// RenameTemplateFiles removes .tmpl extension from template files
func (r *Replacer) RenameTemplateFiles(dirPath string) error {
	var tmplFiles []string

	// Collect all .tmpl files
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".tmpl") {
			tmplFiles = append(tmplFiles, path)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk directory: %w", err)
	}

	// Rename each .tmpl file
	for _, tmplPath := range tmplFiles {
		newPath := strings.TrimSuffix(tmplPath, ".tmpl")
		if err := os.Rename(tmplPath, newPath); err != nil {
			return fmt.Errorf("failed to rename %s to %s: %w", tmplPath, newPath, err)
		}
	}

	return nil
}

// RenameEntityDirectory renames the generic "entity" directory to the actual entity name
func (r *Replacer) RenameEntityDirectory(projectPath string) error {
	entityPath := filepath.Join(projectPath, "internal", "entity")
	newEntityPath := filepath.Join(projectPath, "internal", r.vars.PrimaryEntity)

	// Check if entity directory exists
	if _, err := os.Stat(entityPath); os.IsNotExist(err) {
		return nil // Nothing to rename
	}

	// Rename the directory
	if err := os.Rename(entityPath, newEntityPath); err != nil {
		return fmt.Errorf("failed to rename entity directory: %w", err)
	}

	return nil
}

// GenerateEntityName generates plural and capitalized forms of entity
func GenerateEntityNames(entity string) (capitalized, plural string) {
	// Capitalize first letter
	if len(entity) > 0 {
		capitalized = strings.ToUpper(string(entity[0])) + entity[1:]
	}

	// Simple pluralization (can be enhanced)
	plural = entity
	if strings.HasSuffix(entity, "y") {
		plural = entity[:len(entity)-1] + "ies"
	} else if strings.HasSuffix(entity, "s") || strings.HasSuffix(entity, "sh") || strings.HasSuffix(entity, "ch") {
		plural = entity + "es"
	} else {
		plural = entity + "s"
	}

	return capitalized, plural
}

// CleanGoImports fixes import paths after template processing
func (r *Replacer) CleanGoImports(projectPath string) error {
	return filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Fix import paths that might have been broken during template processing
		contentStr := string(content)

		// Replace any remaining template artifacts
		contentStr = r.fixImportPaths(contentStr)

		if err := os.WriteFile(path, []byte(contentStr), info.Mode()); err != nil {
			return fmt.Errorf("failed to write cleaned file %s: %w", path, err)
		}

		return nil
	})
}

// fixImportPaths cleans up import paths
func (r *Replacer) fixImportPaths(content string) string {
	// Remove any duplicate quotes or malformed import paths
	importRegex := regexp.MustCompile(`"([^"]+/[^"]+)"`)
	return importRegex.ReplaceAllStringFunc(content, func(match string) string {
		// Clean up any template artifacts in import paths
		cleaned := strings.ReplaceAll(match, "//", "/")
		return cleaned
	})
}