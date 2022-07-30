package main

import (
	"context"
	"log"
	"math/rand"

	"github.com/JYisus/PodChaosMonkey/config"
	"github.com/JYisus/PodChaosMonkey/pkg/k8s"
	"github.com/robfig/cron/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	cfg := config.NewConfig()
	clientset := k8s.NewClientset(cfg)
	ctx := context.Background()

	var scheduler *cron.Cron
	if cfg.ScheduleFormat == "cron" {
		scheduler = cron.New()
	} else {
		scheduler = cron.New(cron.WithSeconds())
	}
	_, err := scheduler.AddFunc(cfg.Schedule, func() {
		runningPods, err := clientset.CoreV1().Pods(cfg.Namespace).List(ctx, v1.ListOptions{})
		if err != nil {
			log.Printf("error listing pods on namespace \"%s\": %s", cfg.Namespace, err)
			return
		}
		podToDelete := runningPods.Items[rand.Intn(len(runningPods.Items))]
		err = clientset.CoreV1().Pods(cfg.Namespace).Delete(ctx, podToDelete.GetName(), v1.DeleteOptions{})
		if err != nil {
			log.Printf("error deleting pod \"%s\": %s", podToDelete.GetName(), err)
			return
		}
	})
	if err != nil {
		log.Printf("error adding func to scheduler: %s", err)
		return
	}
	scheduler.Start()
	select {}
}
