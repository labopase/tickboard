package uconfig_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/utilities/uconfig"
)

type TestConfig struct {
	App struct {
		Name string `json:"name" yaml:"name"`
		Port int    `json:"port" yaml:"port"`
	} `json:"app" yaml:"app"`
}

type TestAppConfig struct {
	Name string `json:"name" yaml:"name"`
	Port int    `json:"port" yaml:"port"`
}

// changeDir changes the working directory to dir and registers a cleanup function
// to restore the original directory. Tests using this cannot use t.Parallel().
func changeDir(t *testing.T, dir string) {
	t.Helper()
	old, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current directory: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}
	t.Cleanup(func() {
		os.Chdir(old)
	})
}

func TestBind(t *testing.T) {
	// Not using t.Parallel() here because other tests use changeDir, which alters
	// the process-wide working directory and could cause race conditions.

	dir := t.TempDir()

	// Write test JSON config
	jsonConfig := `{"app": {"name": "test-app", "port": 8080}}`
	err := os.WriteFile(filepath.Join(dir, "myconfig.json"), []byte(jsonConfig), 0644)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	// Write test YAML config
	yamlConfig := `
app:
  name: "test-yaml"
  port: 9090
`
	err = os.WriteFile(filepath.Join(dir, "myconfig.yaml"), []byte(yamlConfig), 0644)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	tests := []struct {
		name      string
		key       string
		path      string
		fname     string
		extension string
		wantName  string
		wantPort  int
		wantErr   bool
	}{
		{
			name:      "Bind JSON entire config",
			key:       "",
			path:      dir,
			fname:     "myconfig",
			extension: "json",
			wantName:  "test-app",
			wantPort:  8080,
			wantErr:   false,
		},
		{
			name:      "Bind YAML entire config",
			key:       "",
			path:      dir,
			fname:     "myconfig",
			extension: "yaml",
			wantName:  "test-yaml",
			wantPort:  9090,
			wantErr:   false,
		},
		{
			name:      "File not found",
			key:       "",
			path:      dir,
			fname:     "notfound",
			extension: "json",
			wantErr:   true,
		},
		{
			name:      "Default values with no config file in cwd",
			key:       "",
			path:      "", // Should default to "./", "config", "json"
			fname:     "",
			extension: "",
			wantErr:   true, // Since "config.json" doesn't exist in current working directory unless created
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := uconfig.Bind[TestConfig](tt.key, tt.path, tt.fname, tt.extension)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.App.Name != tt.wantName {
					t.Errorf("Bind() got Name = %v, want %v", got.App.Name, tt.wantName)
				}
				if got.App.Port != tt.wantPort {
					t.Errorf("Bind() got Port = %v, want %v", got.App.Port, tt.wantPort)
				}
			}
		})
	}
}

func TestBind_Key(t *testing.T) {
	dir := t.TempDir()

	jsonConfig := `{"app": {"name": "test-app", "port": 8080}}`
	err := os.WriteFile(filepath.Join(dir, "myconfig.json"), []byte(jsonConfig), 0644)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	tests := []struct {
		name      string
		key       string
		path      string
		fname     string
		extension string
		wantName  string
		wantPort  int
		wantErr   bool
	}{
		{
			name:      "Bind JSON specific key",
			key:       "app",
			path:      dir,
			fname:     "myconfig",
			extension: "json",
			wantName:  "test-app",
			wantPort:  8080,
			wantErr:   false,
		},
		{
			name:      "Bind JSON missing key",
			key:       "missing",
			path:      dir,
			fname:     "myconfig",
			extension: "json",
			wantName:  "", // Unmarshals into zero values; Viper doesn't error on missing keys
			wantPort:  0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := uconfig.Bind[TestAppConfig](tt.key, tt.path, tt.fname, tt.extension)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Name != tt.wantName {
					t.Errorf("Bind() got Name = %v, want %v", got.Name, tt.wantName)
				}
				if got.Port != tt.wantPort {
					t.Errorf("Bind() got Port = %v, want %v", got.Port, tt.wantPort)
				}
			}
		})
	}
}

func TestBindJSONKey(t *testing.T) {
	dir := t.TempDir()
	changeDir(t, dir)

	jsonConfig := `{"database": {"host": "localhost", "port": 5432}}`
	err := os.WriteFile(filepath.Join(dir, "config.json"), []byte(jsonConfig), 0644)
	if err != nil {
		t.Fatalf("failed to write config.json: %v", err)
	}

	type DBConfig struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}

	tests := []struct {
		name     string
		key      string
		setup    func()
		wantHost string
		wantPort int
		wantErr  bool
	}{
		{
			name:     "Valid Key",
			key:      "database",
			setup:    func() {}, // File already exists
			wantHost: "localhost",
			wantPort: 5432,
			wantErr:  false,
		},
		{
			name: "File Not Found",
			key:  "database",
			setup: func() {
				// Remove config.json to simulate missing file
				os.Remove(filepath.Join(dir, "config.json"))
			},
			wantHost: "",
			wantPort: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, err := uconfig.BindJSONKey[DBConfig](tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("BindJSONKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Host != tt.wantHost {
					t.Errorf("got Host = %v, want %v", got.Host, tt.wantHost)
				}
				if got.Port != tt.wantPort {
					t.Errorf("got Port = %v, want %v", got.Port, tt.wantPort)
				}
			}
		})
	}
}

func TestBindYAMLKey(t *testing.T) {
	dir := t.TempDir()
	changeDir(t, dir)

	yamlConfig := `
database:
  host: "db.local"
  port: 3306
`
	err := os.WriteFile(filepath.Join(dir, "config.yaml"), []byte(yamlConfig), 0644)
	if err != nil {
		t.Fatalf("failed to write config.yaml: %v", err)
	}

	type DBConfig struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	tests := []struct {
		name     string
		key      string
		wantHost string
		wantPort int
		wantErr  bool
	}{
		{
			name:     "Valid Key",
			key:      "database",
			wantHost: "db.local",
			wantPort: 3306,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := uconfig.BindYAMLKey[DBConfig](tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("BindYAMLKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Host != tt.wantHost {
					t.Errorf("got Host = %v, want %v", got.Host, tt.wantHost)
				}
				if got.Port != tt.wantPort {
					t.Errorf("got Port = %v, want %v", got.Port, tt.wantPort)
				}
			}
		})
	}
}
