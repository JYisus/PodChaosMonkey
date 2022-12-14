package config

import (
	"log"
	"os"
	"strings"

	"github.com/JYisus/PodChaosMonkey/pkg/scheduler"
	"github.com/JYisus/PodChaosMonkey/pkg/terminator"

	"gopkg.in/yaml.v3"
)

// Config is the struct that contains the configuration of the PodChaosMonkey.
type Config struct {
	Namespace       string
	Schedule        string
	ScheduleFormat  string
	IsInsideCluster bool
	Blacklist       *terminator.Blacklist
}

const blacklistPath = "blacklist.yml"

// New creates a Config struct from the provided environment variables.
func New() *Config {
	cfg := &Config{
		Namespace:       "workloads",
		Schedule:        "* * * * *",
		ScheduleFormat:  "cron",
		IsInsideCluster: false,
	}
	namespace := os.Getenv("NAMESPACE")
	if namespace != "" {
		cfg.Namespace = namespace
	}

	schedule := os.Getenv("SCHEDULE")
	if schedule != "" {
		cfg.Schedule = schedule
	}

	scheduleFormat := os.Getenv("SCHEDULE_FORMAT")
	if scheduleFormat != "" {
		if !scheduler.IsValidFormat(scheduleFormat) {
			log.Panicf("Invalid schedule format: %s", scheduleFormat)
		}
		cfg.ScheduleFormat = scheduleFormat
	}

	if strings.ToLower(os.Getenv("IS_INSIDE_CLUSTER")) == "true" {
		cfg.IsInsideCluster = true
	}

	cfg.Blacklist = newBlacklist()

	return cfg
}

func newBlacklist() *terminator.Blacklist {
	f, err := os.Open(blacklistPath)
	if err != nil {
		log.Printf("Blacklist configuration file \"%s\" not provided", blacklistPath)
		return nil
	}
	blacklist := &terminator.Blacklist{}
	err = yaml.NewDecoder(f).Decode(blacklist)
	if err != nil {
		panic(err)
	}
	return blacklist
}
