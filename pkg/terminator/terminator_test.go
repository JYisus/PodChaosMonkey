package terminator_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/JYisus/PodChaosMonkey/pkg/terminator"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func generatePods(namespace string, amount int) []runtime.Object {
	pods := make([]runtime.Object, amount)
	for i := 0; i < amount; i++ {
		pods[i] = &corev1.Pod{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Pod",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("pod-%d", i),
				Namespace: namespace,
			},
		}
	}
	return pods
}

func listPods(clientset *fake.Clientset, namespace string) []corev1.Pod {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	return pods.Items
}

func TestKiller_KillRandomPod(t *testing.T) {
	tests := []struct {
		name                 string
		namespace            string
		numberOfPods         int
		expectedNumberOfPods int
		wantError            bool
		errorString          string
	}{
		{
			name:                 "Test successfully delete 1 random pod",
			namespace:            "workloads",
			numberOfPods:         3,
			expectedNumberOfPods: 2,
			wantError:            false,
			errorString:          "",
		},
		{
			name:                 "Test no pod exist in the namespace",
			namespace:            "workloads",
			numberOfPods:         0,
			expectedNumberOfPods: 0,
			wantError:            true,
			errorString:          "no pods running on namespace \"workloads\"",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pods := generatePods(tc.namespace, tc.numberOfPods)
			clientset := fake.NewSimpleClientset(pods...)
			podTerminator := terminator.NewPodTerminator(clientset, nil)
			err := podTerminator.KillRandomPod(context.Background(), tc.namespace)

			if tc.wantError {
				assert.Error(t, err, tc.errorString)
				return
			}

			podsAfterDelete := listPods(clientset, tc.namespace)
			assert.Equal(t, tc.expectedNumberOfPods, len(podsAfterDelete))
		})
	}
}
