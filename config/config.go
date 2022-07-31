package config

import (
	"log"
	"os"
	"strings"

	"github.com/JYisus/PodChaosMonkey/pkg/scheduler"
)

type Config struct {
	Namespace       string
	Schedule        string
	ScheduleFormat  string
	IsInsideCluster bool
}

func NewConfig() *Config {
	cfg := &Config{
		Namespace:       "workloads",
		Schedule:        "* * * * *",
		ScheduleFormat:  "cron",
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

	scheduleFormat := os.Getenv("CM_SCHEDULE_FORMAT")
	if scheduleFormat != "" {
		if !scheduler.IsValidFormat(scheduleFormat) {
			log.Fatalln("Invalid schedule format: ", scheduleFormat)
		}
		cfg.ScheduleFormat = scheduleFormat
	}

	if strings.ToLower(os.Getenv("CM_IS_INSIDE_CLUSTER")) == "true" {
		cfg.IsInsideCluster = true
	}
	return cfg
}
