package chaos

import (
	"context"
	"log"

	"github.com/JYisus/PodChaosMonkey/pkg/scheduler"
	"github.com/JYisus/PodChaosMonkey/pkg/terminator"
)

// Chaos is the struct that contains the scheduler and the terminator.
type Chaos struct {
	cronScheduler scheduler.Scheduler
	podTerminator terminator.Terminator
}

// New returns a new Chaos struct.
func New(cronScheduler scheduler.Scheduler, podTerminator terminator.Terminator) *Chaos {
	return &Chaos{
		cronScheduler: cronScheduler,
		podTerminator: podTerminator,
	}
}

// Run starts the scheduler with the terminator KillRandomPod function.
func (c *Chaos) Run(ctx context.Context, schedule string, namespace string) error {
	return c.cronScheduler.Start(ctx, schedule, func() {
		err := c.podTerminator.KillRandomPod(ctx, namespace)
		if err != nil {
			log.Println(err)
		}
	})
}
