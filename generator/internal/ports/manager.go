package ports

import (
	"crypto/rand"
	"math/big"

	"github.com/kranti/go-template-generator/internal/config"
)

type Manager struct {
	BaseAPIPort      int
	BaseDBPort       int
	BaseRedisPort    int
	Increment        int
	RandomEnabled    bool
	RandomRange      int
}

type Ports struct {
	API   int
	DB    int
	Redis int
}

func NewManager(cfg *config.Config) *Manager {
	return &Manager{
		BaseAPIPort:   cfg.Ports.BaseAPI,
		BaseDBPort:    cfg.Ports.BaseDB,
		BaseRedisPort: cfg.Ports.BaseRedis,
		Increment:     cfg.Ports.Increment,
		RandomEnabled: cfg.Ports.Randomization.Enabled,
		RandomRange:   cfg.Ports.Randomization.Range,
	}
}

// AllocatePorts calculates port numbers based on project index with optional randomization
func (m *Manager) AllocatePorts(projectIndex int) Ports {
	// Calculate base offset from project index
	baseOffset := projectIndex * m.Increment

	// Calculate base ports
	basePorts := Ports{
		API:   m.BaseAPIPort + baseOffset,
		DB:    m.BaseDBPort + baseOffset,
		Redis: m.BaseRedisPort + baseOffset,
	}

	// Apply randomization if enabled
	if m.RandomEnabled && m.RandomRange > 0 {
		return Ports{
			API:   m.applyRandomOffset(basePorts.API),
			DB:    m.applyRandomOffset(basePorts.DB),
			Redis: m.applyRandomOffset(basePorts.Redis),
		}
	}

	return basePorts
}

// applyRandomOffset adds a random offset within the configured range
func (m *Manager) applyRandomOffset(basePort int) int {
	// Generate random offset between -range and +range
	randomOffset := m.generateRandomOffset()

	// Apply offset and ensure port is within reasonable bounds
	newPort := basePort + randomOffset

	// Ensure port is within valid range (1024-65535)
	if newPort < 1024 {
		newPort = 1024 + (newPort % 100)
	} else if newPort > 65535 {
		newPort = 65535 - (newPort % 100)
	}

	return newPort
}

// generateRandomOffset creates a cryptographically secure random offset
func (m *Manager) generateRandomOffset() int {
	// Generate random number between 0 and (2 * range)
	maxOffset := int64(2 * m.RandomRange)
	randomNum, err := rand.Int(rand.Reader, big.NewInt(maxOffset))
	if err != nil {
		// Fallback to no randomization if crypto/rand fails
		return 0
	}

	// Convert to offset range: -range to +range
	offset := int(randomNum.Int64()) - m.RandomRange
	return offset
}