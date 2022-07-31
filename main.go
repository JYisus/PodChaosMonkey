package main

import (
	"context"
	"github.com/JYisus/PodChaosMonkey/pkg/terminator"
	"log"

	"github.com/JYisus/PodChaosMonkey/config"
	"github.com/JYisus/PodChaosMonkey/pkg/chaos"
	"github.com/JYisus/PodChaosMonkey/pkg/k8s"
	"github.com/JYisus/PodChaosMonkey/pkg/scheduler"
)

func main() {
	cfg := config.NewConfig()
	clientset := k8s.NewClientset(cfg)
	ctx := context.Background()
	sch := scheduler.NewCronScheduler(cfg.ScheduleFormat)
	podTerminator := terminator.NewPodTerminator(clientset)
	c := chaos.New(sch, podTerminator)
	if err := c.Run(ctx, cfg.Schedule, cfg.Namespace); err != nil {
		log.Fatalf("Error running chaos: %s", err)
	}
}
