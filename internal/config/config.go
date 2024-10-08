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
	EnvLocal   Env = "local"
	EnvRelease Env = "release"
)

type Config struct {
	Env        Env        `yaml:"env"`
	HttpServer HttpServer `yaml:"httpServer"`
	Cache      Cache      `yaml:"cache"`
}

type HttpServer struct {
	Addr            string        `yaml:"addr"`
	Timeout         time.Duration `yaml:"timeout"`
	IdleTimeout     time.Duration `yaml:"idleTimeout"`
	ShutDownTimeout time.Duration `yaml:"shutDownTimeout"`
}

type Cache struct {
	Cap int `yaml:"cap"`
}

func expandContentWithEnv(input string) string {
	mapping := func(key string) string {
		return os.Getenv(key)
	}

	return os.Expand(input, mapping)
}

func parseYamlWithEnv(file string) (config *Config) {
	content, err := os.ReadFile(file)
	if err != nil {
		return config
	}

	enrichedContent := expandContentWithEnv(string(content))
	err = yaml.Unmarshal([]byte(enrichedContent), &config)
	if err != nil {
		return config
	}

	return config
}

func getConfigPath(env Env) string {
	const prefix = "resources/config"
	switch env {
	case EnvLocal:
		return prefix + "/local.yaml"
	case EnvRelease:
		return prefix + "/release.yaml"
	}
	return ""
}

func MustGet(env Env) *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(ErrFailedToLoadEnv)
	}

	path := getConfigPath(env)
	config := parseYamlWithEnv(path)
	if config == nil {
		log.Fatal(ErrParsingYaml)
	}

	return config
}
