package terminator

import (
	"context"
	"fmt"
	"math/rand"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//go:generate mockgen -source=terminator.go -destination=../mocks/terminator_mock.go -package=mocks

type Terminator interface {
	KillRandomPod(ctx context.Context, namespace string) error
}

type PodTerminator struct {
	k8sClient kubernetes.Interface
}

func NewPodTerminator(k8sClient kubernetes.Interface) *PodTerminator {
	return &PodTerminator{
		k8sClient: k8sClient,
	}
}

func (t *PodTerminator) KillRandomPod(ctx context.Context, namespace string) error {
	runningPods, err := t.k8sClient.CoreV1().Pods(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error listing pods on namespace \"%s\": %s", namespace, err)
	}

	if len(runningPods.Items) == 0 {
		return fmt.Errorf("no pods running on namespace \"%s\"", namespace)
	}

	podToDelete := runningPods.Items[rand.Intn(len(runningPods.Items))]

	err = t.k8sClient.CoreV1().Pods(namespace).Delete(ctx, podToDelete.GetName(), v1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("error deleting pod \"%s\" on namespace \"%s\": %s", podToDelete.GetName(), namespace, err)
	}
	return nil
}
