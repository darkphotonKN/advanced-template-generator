package registry

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Project struct {
	Name      string    `json:"name"`
	Index     int       `json:"index"`
	APIPort   int       `json:"api_port"`
	DBPort    int       `json:"db_port"`
	RedisPort int       `json:"redis_port"`
	Entity    string    `json:"entity"`
	CreatedAt time.Time `json:"created_at"`
}

type Registry struct {
	Projects  []Project `json:"projects"`
	NextIndex int       `json:"next_index"`
}

type Manager struct {
	registryPath string
}

func NewManager(registryPath string) *Manager {
	return &Manager{registryPath: registryPath}
}

func (m *Manager) Load() (*Registry, error) {
	// Ensure the directory exists
	dir := filepath.Dir(m.registryPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create registry directory: %w", err)
	}

	// If file doesn't exist, return empty registry
	if _, err := os.Stat(m.registryPath); os.IsNotExist(err) {
		return &Registry{
			Projects:  []Project{},
			NextIndex: 1, // Start at 1 so first project gets offset ports
		}, nil
	}

	data, err := os.ReadFile(m.registryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read registry file: %w", err)
	}

	var registry Registry
	if err := json.Unmarshal(data, &registry); err != nil {
		return nil, fmt.Errorf("failed to parse registry file: %w", err)
	}

	return &registry, nil
}

func (m *Manager) Save(registry *Registry) error {
	data, err := json.MarshalIndent(registry, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal registry: %w", err)
	}

	if err := os.WriteFile(m.registryPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write registry file: %w", err)
	}

	return nil
}

func (m *Manager) AddProject(name, entity string, apiPort, dbPort, redisPort int) error {
	registry, err := m.Load()
	if err != nil {
		return err
	}

	// Check if project already exists
	for _, project := range registry.Projects {
		if project.Name == name {
			return fmt.Errorf("project '%s' already exists", name)
		}
	}

	// Add new project
	project := Project{
		Name:      name,
		Index:     registry.NextIndex,
		APIPort:   apiPort,
		DBPort:    dbPort,
		RedisPort: redisPort,
		Entity:    entity,
		CreatedAt: time.Now(),
	}

	registry.Projects = append(registry.Projects, project)
	registry.NextIndex++

	return m.Save(registry)
}

func (m *Manager) GetNextIndex() (int, error) {
	registry, err := m.Load()
	if err != nil {
		return 0, err
	}
	return registry.NextIndex, nil
}

func (m *Manager) List() ([]Project, error) {
	registry, err := m.Load()
	if err != nil {
		return nil, err
	}
	return registry.Projects, nil
}

func (m *Manager) ProjectExists(name string) (bool, error) {
	registry, err := m.Load()
	if err != nil {
		return false, err
	}

	for _, project := range registry.Projects {
		if project.Name == name {
			return true, nil
		}
	}
	return false, nil
}