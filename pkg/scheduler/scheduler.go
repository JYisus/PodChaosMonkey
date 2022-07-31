package scheduler

import (
	"context"
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

//go:generate mockgen -source=scheduler.go -destination=../mocks/scheduler_mock.go -package=mocks

type Scheduler interface {
	Start(ctx context.Context, schedule string, task func()) error
}

type CronScheduler struct {
	cron *cron.Cron
}

func NewCronScheduler(scheduleFormat string) *CronScheduler {
	return &CronScheduler{
		cron: getCronWithFormat(scheduleFormat),
	}
}

func getCronWithFormat(scheduleFormat string) *cron.Cron {
	switch {
	case scheduleFormat == CronFormat:
		return cron.New()
	case scheduleFormat == CronSecondsFormat:
		return cron.New(cron.WithSeconds())
	default:
		log.Fatalln("Unknown schedule format")
	}
	return nil
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
