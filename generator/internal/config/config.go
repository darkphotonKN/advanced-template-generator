package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		User        string `yaml:"user"`
		Password    string `yaml:"password"`
		NamePattern string `yaml:"name_pattern"`
	} `yaml:"database"`

	Ports struct {
		BaseAPI       int `yaml:"base_api"`
		BaseDB        int `yaml:"base_db"`
		BaseRedis     int `yaml:"base_redis"`
		BaseFrontend  int `yaml:"base_frontend"`
		Increment     int `yaml:"increment"`
		Randomization struct {
			Enabled bool `yaml:"enabled"`
			Range   int  `yaml:"range"`
		} `yaml:"randomization"`
	} `yaml:"ports"`

	Defaults struct {
		ModulePrefix    string `yaml:"module_prefix"`
		IncludeAuth     bool   `yaml:"include_auth"`
		IncludeRedis    bool   `yaml:"include_redis"`
		IncludeS3       bool   `yaml:"include_s3"`
		IncludeFrontend bool   `yaml:"include_frontend"`
		PrimaryEntity   string `yaml:"primary_entity"`
	} `yaml:"defaults"`

	ProjectsRegistry string `yaml:"projects_registry"`

	Git struct {
		InitialCommitMessage string `yaml:"initial_commit_message"`
	} `yaml:"git"`

	Features struct {
		Auth struct {
			Enabled     bool   `yaml:"enabled"`
			Description string `yaml:"description"`
		} `yaml:"auth"`
		S3 struct {
			Enabled     bool   `yaml:"enabled"`
			Description string `yaml:"description"`
		} `yaml:"s3"`
		Redis struct {
			Enabled     bool   `yaml:"enabled"`
			Description string `yaml:"description"`
		} `yaml:"redis"`
		Frontend struct {
			Enabled     bool   `yaml:"enabled"`
			Description string `yaml:"description"`
		} `yaml:"frontend"`
	} `yaml:"features"`
}

func LoadConfig(path string) (*Config, error) {
	// Check if path is relative or absolute
	if !filepath.IsAbs(path) {
		// If relative, look in the generator directory
		execPath, err := os.Executable()
		if err != nil {
			return nil, fmt.Errorf("failed to get executable path: %w", err)
		}
		path = filepath.Join(filepath.Dir(execPath), "..", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Expand home directory in registry path
	if config.ProjectsRegistry != "" && config.ProjectsRegistry[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get home directory: %w", err)
		}
		config.ProjectsRegistry = filepath.Join(home, config.ProjectsRegistry[1:])
	}

	return &config, nil
}