package config

import (
	"os"
	"strings"
)

type Config struct {
	Namespace       string
	Schedule        string
	IsInsideCluster bool
}

func NewConfig() *Config {
	cfg := &Config{
		Namespace:       "workloads",
		Schedule:        "*/5 * * * * *",
		IsInsideCluster: false,
	}
	namespace := os.Getenv("CM_NAMESPACE")
	if namespace != "" {
		cfg.Namespace = namespace
	}

	schedule := os.Getenv("CM_SCHEDULE")
	if schedule != "" {
		cfg.Schedule = schedule
	}
	if strings.ToLower(os.Getenv("CM_IS_INSIDE_CLUSTER")) == "true" {
		cfg.IsInsideCluster = true
	}
	return cfg
}
