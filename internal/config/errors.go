package config

import "errors"

var (
	ErrFailedToLoadEnv = errors.New("failed to load .env")
	ErrParsingYaml     = errors.New("error parsing yaml")
)
