package scheduler

import (
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
)

//go:generate mockgen -source=scheduler.go -destination=../mocks/scheduler_mock.go -package=mocks

// Scheduler is an interface that defines the methods that a scheduler must implement.
type Scheduler interface {
	Start(ctx context.Context, schedule string, task func()) error
}

// CronScheduler is a struct that manage the cron scheduler.
type CronScheduler struct {
	cron *cron.Cron
}

// NewCronScheduler creates a new CronScheduler with the schedule format provided.
func NewCronScheduler(scheduleFormat string) (*CronScheduler, error) {
	c, err := getCronWithFormat(scheduleFormat)
	if err != nil {
		return nil, err
	}
	return &CronScheduler{
		cron: c,
	}, nil
}

func getCronWithFormat(scheduleFormat string) (*cron.Cron, error) {
	switch {
	case scheduleFormat == CronFormat:
		return cron.New(), nil
	case scheduleFormat == CronSecondsFormat:
		return cron.New(cron.WithSeconds()), nil
	default:
		fmt.Println("schedule format not supported")
		return nil, fmt.Errorf("unknown schedule format: %s", scheduleFormat)
	}
}

func (s *CronScheduler) Start(ctx context.Context, schedule string, task func()) error {
	_, err := s.cron.AddFunc(schedule, task)
	if err != nil {
		return fmt.Errorf("error adding task to scheduler: %s", err)
	}

	s.cron.Start()
	select {
	case <-ctx.Done():
		s.cron.Stop()
		return nil
	}
}
