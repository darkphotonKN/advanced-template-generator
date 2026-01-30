package main

import (
	"fmt"
	"os"

	"github.com/darkphotonKN/go-template-generator/internal/config"
	"github.com/darkphotonKN/go-template-generator/internal/ddd"
	"github.com/spf13/cobra"
)

var (
	// Flags
	entity       string
	noAuth       bool
	withS3       bool
	withFrontend bool
	description  string
)

var rootCmd = &cobra.Command{
	Use:   "go-gen",
	Short: "Go Template Generator - Create production-ready Go projects",
	Long:  `A powerful Go project generator that creates production-ready applications with consistent structure, best practices, and pre-configured development environment.`,
}

var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "Create a new DDD API project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Load configuration
		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			os.Exit(1)
		}

		// Override defaults with flags
		if entity != "" {
			cfg.Defaults.PrimaryEntity = entity
		}
		if noAuth {
			cfg.Features.Auth.Enabled = false
		}
		if withS3 {
			cfg.Features.S3.Enabled = true
		}
		if withFrontend {
			cfg.Features.Frontend.Enabled = true
		}
		if description == "" {
			description = fmt.Sprintf("DDD API for %s management", cfg.Defaults.PrimaryEntity)
		}

		// Create generator options
		opts := &ddd.GeneratorOptions{
			ProjectName:        projectName,
			Entity:            cfg.Defaults.PrimaryEntity,
			IncludeAuth:       cfg.Features.Auth.Enabled,
			IncludeS3:         cfg.Features.S3.Enabled,
			IncludeRedis:      cfg.Features.Redis.Enabled,
			IncludeFrontend:   cfg.Features.Frontend.Enabled,
			ProjectDescription: description,
			Config:            cfg,
		}

		// Generate the project
		generator := ddd.NewGenerator(opts)
		if err := generator.Generate(); err != nil {
			fmt.Printf("Error generating project: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\n✅ Project '%s' created successfully!\n\n", projectName)
		fmt.Printf("Next steps:\n")
		if opts.IncludeFrontend {
			fmt.Printf("  cd %s/%s-server\n", projectName, projectName)
		} else {
			fmt.Printf("  cd %s\n", projectName)
		}
		fmt.Printf("  cp .env.example .env\n")
		fmt.Printf("  make docker-up\n")
		fmt.Printf("  make migrate-up\n")
		fmt.Printf("  make dev\n\n")
		fmt.Printf("Your API will be running at http://localhost:%d\n", opts.APIPort)
		if opts.IncludeFrontend {
			fmt.Printf("Your frontend will be running at http://localhost:%d\n", opts.FrontendPort)
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all generated projects",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			os.Exit(1)
		}

		registry := ddd.NewProjectRegistry(cfg.ProjectsRegistry)
		projects, err := registry.List()
		if err != nil {
			fmt.Printf("Error listing projects: %v\n", err)
			os.Exit(1)
		}

		if len(projects) == 0 {
			fmt.Println("No projects generated yet.")
			return
		}

		fmt.Println("Generated projects:")
		fmt.Println("==================")
		for _, p := range projects {
			fmt.Printf("  %s - API: %d, DB: %d, Redis: %d (created: %s)\n",
				p.Name, p.APIPort, p.DBPort, p.RedisPort, p.CreatedAt.Format("2006-01-02"))
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&entity, "entity", "e", "", "Primary entity name (default: item)")
	createCmd.Flags().BoolVar(&noAuth, "no-auth", false, "Generate without authentication")
	createCmd.Flags().BoolVar(&withS3, "with-s3", false, "Include S3 file upload support")
	createCmd.Flags().BoolVar(&withFrontend, "with-frontend", false, "Include Next.js frontend")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "Project description for CLAUDE.md")

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}