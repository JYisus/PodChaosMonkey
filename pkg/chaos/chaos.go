package chaos

import (
	"context"
	"github.com/JYisus/PodChaosMonkey/pkg/scheduler"
	"github.com/JYisus/PodChaosMonkey/pkg/terminator"
	"log"
)

type Chaos struct {
	cronScheduler scheduler.Scheduler
	podTerminator terminator.Terminator
}

func New(cronScheduler scheduler.Scheduler, podTerminator terminator.Terminator) *Chaos {
	return &Chaos{
		cronScheduler: cronScheduler,
		podTerminator: podTerminator,
	}
}

func (c *Chaos) Run(ctx context.Context, schedule string, namespace string) error {
	return c.cronScheduler.Start(ctx, schedule, func() {
		err := c.podTerminator.KillRandomPod(ctx, namespace)
		if err != nil {
			log.Println(err)
		}
	})
}
