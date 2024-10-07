package config

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Env string

const (
	EnvLocal = "local"
)

type Config struct {
	Env        Env        `yaml:"env"`
	HttpServer HttpServer `yaml:"httpServer"`
}

type HttpServer struct {
	Addr            string        `yaml:"addr"`
	Port            int           `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	IdleTimeout     time.Duration `yaml:"idleTimeout"`
	ShutDownTimeout time.Duration `yaml:"shutDownTimeout"`
}

func expandContentWithEnvVars(input string) string {
	mapping := func(key string) string {
		return os.Getenv(key)
	}

	return os.Expand(input, mapping)
}

func parseYamlWithEnv[T any](file string) *T {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	contentWithEnv := expandContentWithEnvVars(string(content))

	var config T
	err = yaml.Unmarshal([]byte(contentWithEnv), &config)
	if err != nil {
		return nil
	}

	return &config
}

func MustGet[T any](path string) *T {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(ErrFailedToLoadEnv)
	}

	config := parseYamlWithEnv[T](path)
	if config == nil {
		log.Fatal(ErrParsingYaml)
	}

	return config
}
