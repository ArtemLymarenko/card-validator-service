package main

import (
	"card-validator-service/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustGet(config.EnvLocal)
	fmt.Println(cfg)
}
