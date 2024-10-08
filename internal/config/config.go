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
	EnvLocal Env = "local"
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

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Dialect  string `yaml:"dialect"`
	Port     int    `yaml:"port"`
	PoolMin  int    `yaml:"poolMin"`
	PoolMax  int    `yaml:"poolMax"`
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
