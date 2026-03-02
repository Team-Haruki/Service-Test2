package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config is the top-level configuration for Haruki-Command-Parser.
type Config struct {
	Server      ServerConfig      `yaml:"server"`
	ServiceAPI  ServiceAPIConfig  `yaml:"service_api"`
	HarukiCloud HarukiCloudConfig `yaml:"haruki_cloud"`
	Log         LogConfig         `yaml:"log"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// ServiceAPIConfig holds the connection info for Part2 (Haruki-Service-API).
// Part1 uses this to forward parsed commands and receive rendered images.
type ServiceAPIConfig struct {
	BaseURL    string        `yaml:"base_url"` // e.g. "http://localhost:24000"
	Timeout    string        `yaml:"timeout"`  // e.g. "30s" (parsed into TimeoutDur)
	TimeoutDur time.Duration `yaml:"-"`
}

type HarukiCloudConfig struct {
	SekaiDB DatabaseConfig `yaml:"sekai_db"`
	Region  string         `yaml:"region"`
	// CacheRefreshInterval controls how often character nicknames are
	// reloaded from the database (e.g. "6h", "30m", "0" to disable).
	// Stored as string, parsed by Load() into CacheRefreshIntervalDur.
	CacheRefreshInterval    string        `yaml:"cache_refresh_interval"`
	CacheRefreshIntervalDur time.Duration `yaml:"-"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	DSN      string `yaml:"dsn"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

type LogConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

// Load loads configuration from the given YAML file path.
func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer f.Close()

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	// Parse cache refresh interval string (e.g. "6h", "30m")
	raw := cfg.HarukiCloud.CacheRefreshInterval
	if raw != "" && raw != "0" {
		dur, err := time.ParseDuration(raw)
		if err != nil {
			return nil, fmt.Errorf("invalid cache_refresh_interval %q: %w", raw, err)
		}
		cfg.HarukiCloud.CacheRefreshIntervalDur = dur
	}

	// Parse ServiceAPI timeout (default 60s if not set)
	if cfg.ServiceAPI.Timeout == "" {
		cfg.ServiceAPI.TimeoutDur = 60 * time.Second
	} else {
		dur, err := time.ParseDuration(cfg.ServiceAPI.Timeout)
		if err != nil {
			return nil, fmt.Errorf("invalid service_api.timeout %q: %w", cfg.ServiceAPI.Timeout, err)
		}
		cfg.ServiceAPI.TimeoutDur = dur
	}

	return &cfg, nil
}

// BuildDSN constructs a DSN string from structured database config fields.
func BuildDSN(db DatabaseConfig) (string, error) {
	switch db.Driver {
	case "postgres", "postgresql":
		if db.Host == "" || db.Database == "" || db.User == "" {
			return "", fmt.Errorf("postgres config requires host, database, user")
		}
		port := db.Port
		if port == 0 {
			port = 5432
		}
		sslMode := db.SSLMode
		if sslMode == "" {
			sslMode = "disable"
		}
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			db.Host, port, db.User, db.Password, db.Database, sslMode), nil
	case "sqlite", "sqlite3":
		if db.Database == "" {
			return "", fmt.Errorf("sqlite config requires database path")
		}
		return db.Database, nil
	default:
		return "", fmt.Errorf("unsupported driver: %s", db.Driver)
	}
}
