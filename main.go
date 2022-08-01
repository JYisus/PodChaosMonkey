package main

import (
	"context"
	"log"

	"github.com/JYisus/PodChaosMonkey/pkg/chaos"
	"github.com/JYisus/PodChaosMonkey/pkg/config"
	"github.com/JYisus/PodChaosMonkey/pkg/k8s"
	"github.com/JYisus/PodChaosMonkey/pkg/scheduler"
	"github.com/JYisus/PodChaosMonkey/pkg/terminator"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()
	clientset := k8s.NewClientset(cfg)
	podTerminator := terminator.NewPodTerminator(clientset, cfg.Blacklist)
	sch, err := scheduler.NewCronScheduler(cfg.ScheduleFormat)
	if err != nil {
		log.Fatalf("Fatal error initializing scheduler: %s", err)
	}

	c := chaos.New(sch, podTerminator)
	if err := c.Run(ctx, cfg.Schedule, cfg.Namespace); err != nil {
		log.Fatalf("Error running chaos: %s", err)
	}
}
