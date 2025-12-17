package config

import (
	"flag"
	"os"
	"testing"
)

func TestLoadYamlConfigFromFile(t *testing.T) {
	// Setup env var
	os.Setenv("EVN_TEST_DSN", "postgres://user:pass@localhost:5432/db")
	defer os.Unsetenv("EVN_TEST_DSN")

	service := NewConfigService()

	// 1. Test relative path (No Env File)
	t.Run("Relative Path No Env File", func(t *testing.T) {
		cfg, err := service.LoadYamlConfigFromFile("./test_config.yaml", "")
		if err != nil {
			t.Fatalf("Failed to load config: %v", err)
		}
		if cfg.App.Name != "test-app" {
			t.Errorf("Expected app name test-app, got %s", cfg.App.Name)
		}
		// Expects OS env var which was set at top of test
		if cfg.Db.Dsn != "postgres://user:pass@localhost:5432/db" {
			t.Errorf("Expected DSN from OS env, got %s", cfg.Db.Dsn)
		}
	})

	// 2. Test Env File Override
	t.Run("Env File Override", func(t *testing.T) {
		cfg, err := service.LoadYamlConfigFromFile("./test_config.yaml", "./test.env")
		if err != nil {
			t.Fatalf("Failed to load config with env file: %v", err)
		}
		// test.env has "postgres://envfile:pass@localhost:5432/db"
		expected := "postgres://envfile:pass@localhost:5432/db"
		if cfg.Db.Dsn != expected {
			t.Errorf("Expected DSN from Env File (%s), got %s", expected, cfg.Db.Dsn)
		}
	})

	// 3. Test absolute path (Skipped)
	t.Run("Absolute Path", func(t *testing.T) {
		// Skipped due to platform specific path handling
	})

	// 4. Test Invalid Path
	t.Run("Invalid Path", func(t *testing.T) {
		_, err := service.LoadYamlConfigFromFile("invalid/path.yaml", "")
		if err == nil {
			t.Error("Expected error for invalid path, got nil")
		}
	})

	// 5. Test Arg Expansion
	t.Run("Arg Expansion", func(t *testing.T) {
		// Define flag
		flagName := "cache-user"
		flagVal := "redis-user"
		// Check if flag already exists to avoid redefinition panic in tests
		if flag.Lookup(flagName) == nil {
			flag.String(flagName, flagVal, "cache user")
		} else {
			flag.Set(flagName, flagVal)
		}

		argYaml := `
cache:
 user: "$ARG_CACHE_USER"
`
		tmpfile := "arg_test.yaml"
		os.WriteFile(tmpfile, []byte(argYaml), 0644)
		defer os.Remove(tmpfile)

		cfg, err := service.LoadYamlConfigFromFile("./arg_test.yaml", "")
		if err != nil {
			t.Fatalf("Failed to load arg config: %v", err)
		}

		if cfg.Cache.User != flagVal {
			t.Errorf("Expected cache user %s, got %s", flagVal, cfg.Cache.User)
		}
	})
}
