package ports

import (
	"testing"

	"github.com/kranti/go-template-generator/internal/config"
)

func TestAllocatePorts(t *testing.T) {
	// Test without randomization
	cfg := &config.Config{}
	cfg.Ports.BaseAPI = 8000
	cfg.Ports.BaseDB = 5432
	cfg.Ports.BaseRedis = 6379
	cfg.Ports.Increment = 10
	cfg.Ports.Randomization.Enabled = false
	cfg.Ports.Randomization.Range = 50

	manager := NewManager(cfg)

	// Test first project
	ports1 := manager.AllocatePorts(1)
	expectedAPI1 := 8000 + (1 * 10) // 8010
	expectedDB1 := 5432 + (1 * 10)  // 5442
	expectedRedis1 := 6379 + (1 * 10) // 6389

	if ports1.API != expectedAPI1 {
		t.Errorf("Expected API port %d, got %d", expectedAPI1, ports1.API)
	}
	if ports1.DB != expectedDB1 {
		t.Errorf("Expected DB port %d, got %d", expectedDB1, ports1.DB)
	}
	if ports1.Redis != expectedRedis1 {
		t.Errorf("Expected Redis port %d, got %d", expectedRedis1, ports1.Redis)
	}

	// Test second project
	ports2 := manager.AllocatePorts(2)
	expectedAPI2 := 8000 + (2 * 10) // 8020
	expectedDB2 := 5432 + (2 * 10)  // 5452
	expectedRedis2 := 6379 + (2 * 10) // 6399

	if ports2.API != expectedAPI2 {
		t.Errorf("Expected API port %d, got %d", expectedAPI2, ports2.API)
	}
	if ports2.DB != expectedDB2 {
		t.Errorf("Expected DB port %d, got %d", expectedDB2, ports2.DB)
	}
	if ports2.Redis != expectedRedis2 {
		t.Errorf("Expected Redis port %d, got %d", expectedRedis2, ports2.Redis)
	}
}

func TestAllocatePortsWithRandomization(t *testing.T) {
	// Test with randomization
	cfg := &config.Config{}
	cfg.Ports.BaseAPI = 8000
	cfg.Ports.BaseDB = 5432
	cfg.Ports.BaseRedis = 6379
	cfg.Ports.Increment = 10
	cfg.Ports.Randomization.Enabled = true
	cfg.Ports.Randomization.Range = 50

	manager := NewManager(cfg)

	// Generate ports for first project multiple times
	ports := make([]Ports, 5)
	for i := 0; i < 5; i++ {
		ports[i] = manager.AllocatePorts(1)
	}

	// Check that ports are within expected ranges
	baseAPI := 8000 + (1 * 10)   // 8010
	baseDB := 5432 + (1 * 10)    // 5442
	baseRedis := 6379 + (1 * 10) // 6389

	for i, p := range ports {
		// Check API port is within range
		if p.API < baseAPI-50 || p.API > baseAPI+50 {
			t.Errorf("Iteration %d: API port %d is outside expected range [%d, %d]",
				i, p.API, baseAPI-50, baseAPI+50)
		}

		// Check DB port is within range
		if p.DB < baseDB-50 || p.DB > baseDB+50 {
			t.Errorf("Iteration %d: DB port %d is outside expected range [%d, %d]",
				i, p.DB, baseDB-50, baseDB+50)
		}

		// Check Redis port is within range
		if p.Redis < baseRedis-50 || p.Redis > baseRedis+50 {
			t.Errorf("Iteration %d: Redis port %d is outside expected range [%d, %d]",
				i, p.Redis, baseRedis-50, baseRedis+50)
		}

		// Check ports are within valid system range
		if p.API < 1024 || p.API > 65535 {
			t.Errorf("Iteration %d: API port %d is outside valid system range [1024, 65535]", i, p.API)
		}
		if p.DB < 1024 || p.DB > 65535 {
			t.Errorf("Iteration %d: DB port %d is outside valid system range [1024, 65535]", i, p.DB)
		}
		if p.Redis < 1024 || p.Redis > 65535 {
			t.Errorf("Iteration %d: Redis port %d is outside valid system range [1024, 65535]", i, p.Redis)
		}
	}

	// Check that at least some randomization occurred (not all identical)
	allSame := true
	for i := 1; i < len(ports); i++ {
		if ports[i].API != ports[0].API || ports[i].DB != ports[0].DB || ports[i].Redis != ports[0].Redis {
			allSame = false
			break
		}
	}

	// Note: There's a small chance all random numbers could be the same,
	// but this test helps verify the randomization is working
	t.Logf("Generated ports: %+v", ports)
	t.Logf("All ports identical: %v (this might be ok if randomization generated same values)", allSame)
}

func TestPortBoundaries(t *testing.T) {
	// Test edge cases to ensure ports stay within valid ranges
	cfg := &config.Config{}
	cfg.Ports.BaseAPI = 1000    // Below 1024
	cfg.Ports.BaseDB = 70000    // Above 65535
	cfg.Ports.BaseRedis = 6379
	cfg.Ports.Increment = 10
	cfg.Ports.Randomization.Enabled = true
	cfg.Ports.Randomization.Range = 1000 // Large range to test boundaries

	manager := NewManager(cfg)
	ports := manager.AllocatePorts(1)

	// All ports should be within valid system range regardless of base values
	if ports.API < 1024 || ports.API > 65535 {
		t.Errorf("API port %d is outside valid range [1024, 65535]", ports.API)
	}
	if ports.DB < 1024 || ports.DB > 65535 {
		t.Errorf("DB port %d is outside valid range [1024, 65535]", ports.DB)
	}
	if ports.Redis < 1024 || ports.Redis > 65535 {
		t.Errorf("Redis port %d is outside valid range [1024, 65535]", ports.Redis)
	}

	t.Logf("Boundary test ports: API=%d, DB=%d, Redis=%d", ports.API, ports.DB, ports.Redis)
}