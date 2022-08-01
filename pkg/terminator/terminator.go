package terminator

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//go:generate mockgen -source=terminator.go -destination=../mocks/terminator_mock.go -package=mocks

type Terminator interface {
	KillRandomPod(ctx context.Context, namespace string) error
}

type PodTerminator struct {
	k8sClient       kubernetes.Interface
	listPodsOptions metav1.ListOptions
}

func NewPodTerminator(k8sClient kubernetes.Interface, blacklist *Blacklist) *PodTerminator {
	return &PodTerminator{
		k8sClient:       k8sClient,
		listPodsOptions: getListOptions(blacklist),
	}
}

func (t *PodTerminator) KillRandomPod(ctx context.Context, namespace string) error {
	runningPods, err := t.k8sClient.CoreV1().Pods(namespace).List(ctx, t.listPodsOptions)
	if err != nil {
		return fmt.Errorf("error listing pods on namespace \"%s\": %s", namespace, err)
	}

	if len(runningPods.Items) == 0 {
		return fmt.Errorf("no pods running on namespace \"%s\"", namespace)
	}

	podToDelete := runningPods.Items[rand.Intn(len(runningPods.Items))]

	err = t.k8sClient.CoreV1().Pods(namespace).Delete(ctx, podToDelete.GetName(), metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("error deleting pod \"%s\" on namespace \"%s\": %s", podToDelete.GetName(), namespace, err)
	}
	log.Printf("Pod \"%s\" deleted on namespace \"%s\"", podToDelete.GetName(), podToDelete.GetNamespace())

	return nil
}
