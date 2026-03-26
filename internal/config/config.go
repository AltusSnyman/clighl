package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	configDirName  = ".clighl"
	configFileName = "config.yaml"
)

type Config struct {
	LocationID  string `yaml:"location_id"`
	AccessToken string `yaml:"access_token"`
	APIVersion  string `yaml:"api_version"`
}

// ConfigDir returns the path to ~/.clighl/
func ConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}
	return filepath.Join(home, configDirName), nil
}

// ConfigPath returns the full path to the config file.
func ConfigPath() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, configFileName), nil
}

// Load reads the config from ~/.clighl/config.yaml.
// Environment variables CLIGHL_LOCATION_ID and CLIGHL_ACCESS_TOKEN override file values
// and work even without a config file.
func Load() (*Config, error) {
	var cfg Config

	// Try loading from file (not fatal if missing)
	path, err := ConfigPath()
	if err == nil {
		data, err := os.ReadFile(path)
		if err == nil {
			_ = yaml.Unmarshal(data, &cfg)
		}
	}

	// Environment variable overrides (work without config file)
	if v := os.Getenv("CLIGHL_LOCATION_ID"); v != "" {
		cfg.LocationID = v
	}
	if v := os.Getenv("CLIGHL_ACCESS_TOKEN"); v != "" {
		cfg.AccessToken = v
	}

	if cfg.APIVersion == "" {
		cfg.APIVersion = "2021-07-28"
	}

	if cfg.LocationID == "" || cfg.AccessToken == "" {
		return nil, fmt.Errorf("not authenticated. Run `clighl auth` to set up credentials")
	}

	return &cfg, nil
}

// Save writes the config to ~/.clighl/config.yaml.
func Save(cfg *Config) error {
	if cfg.APIVersion == "" {
		cfg.APIVersion = "2021-07-28"
	}

	dir, err := ConfigDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("could not create config directory: %w", err)
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("could not marshal config: %w", err)
	}

	path := filepath.Join(dir, configFileName)
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("could not write config: %w", err)
	}

	return nil
}

// Delete removes the config file.
func Delete() error {
	path, err := ConfigPath()
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("could not remove config: %w", err)
	}
	return nil
}

// Exists checks if the config file exists and is valid.
func Exists() bool {
	cfg, err := Load()
	return err == nil && cfg != nil
}

// MaskToken returns a masked version of the token for display.
func MaskToken(token string) string {
	if len(token) <= 8 {
		return "****"
	}
	return token[:4] + strings.Repeat("*", len(token)-8) + token[len(token)-4:]
}
