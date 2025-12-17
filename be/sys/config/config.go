package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type ConfigService interface {
	// GetAppDir return app dir
	GetAppDir() (string, error)
	/*
		LoadYamlConfigFromFile load yaml config from file
		return config struct and error
		if yayamlFilePath start with "./" combine it with app dir for full filename
		if yayamlFilePath start with "/" that mean it is absolute path
		else return error invalid path
		evnFilePath is environment variable file path
		function will load evnFilePath first then load yamlFilePath and override yamlFilePath value with evnFilePath value

		if any value start with $ARG_ it will load from argument by using flag package
	*/
	LoadYamlConfigFromFile(yamlFilePath, evnFilePath string) (Config, error)
}

type configService struct{}

func NewConfigService() ConfigService {
	return &configService{}
}

func (s *configService) GetAppDir() (string, error) {
	return os.Getwd()
}

func (s *configService) LoadYamlConfigFromFile(yamlFilePath, evnFilePath string) (Config, error) {
	var cfg Config

	// 1. Resolve and Load Env File
	envMap := make(map[string]string)
	if evnFilePath != "" {
		var finalEnvPath string
		if strings.HasPrefix(evnFilePath, "/") {
			finalEnvPath = evnFilePath
		} else if strings.HasPrefix(evnFilePath, "./") || strings.HasPrefix(evnFilePath, "../") {
			appDir, err := s.GetAppDir()
			if err != nil {
				return cfg, err
			}
			finalEnvPath = filepath.Join(appDir, evnFilePath)
		} else {
			return cfg, fmt.Errorf("invalid env path: %s", evnFilePath)
		}

		var err error
		envMap, err = godotenv.Read(finalEnvPath)
		if err != nil {
			return cfg, fmt.Errorf("failed to load env file: %w", err)
		}
	}

	// 2. Resolve Yaml Path
	var finalPath string
	if strings.HasPrefix(yamlFilePath, "/") {
		finalPath = yamlFilePath
	} else if strings.HasPrefix(yamlFilePath, "./") || strings.HasPrefix(yamlFilePath, "../") {
		appDir, err := s.GetAppDir()
		if err != nil {
			return cfg, err
		}
		finalPath = filepath.Join(appDir, yamlFilePath)
	} else {
		return cfg, fmt.Errorf("invalid path: %s", yamlFilePath)
	}

	data, err := os.ReadFile(finalPath)
	if err != nil {
		return cfg, err
	}

	// 3. Expand environment variables
	expandedData := os.Expand(string(data), func(key string) string {
		if strings.HasPrefix(key, "EVN_") {
			// Check env file first
			if val, ok := envMap[key]; ok {
				return val
			}
			// Fallback to OS env
			return os.Getenv(key)
		}
		if strings.HasPrefix(key, "ARG_") {
			argName := strings.TrimPrefix(key, "ARG_")
			argName = strings.ToLower(argName)
			argName = strings.ReplaceAll(argName, "_", "-")
			f := flag.Lookup(argName)
			if f != nil {
				return f.Value.String()
			}
			return ""
		}
		return "$" + key
	})

	err = yaml.Unmarshal([]byte(expandedData), &cfg)
	if err != nil {
		return cfg, err
	}

	cfg.AppDir = filepath.Dir(finalPath)

	return cfg, nil
}
