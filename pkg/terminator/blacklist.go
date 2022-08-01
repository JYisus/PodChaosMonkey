package terminator

import (
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Blacklist struct {
	Labels         map[string]string `yaml:"labels"`
	FieldSelectors map[string]string `yaml:"fieldSelectors"`
}

func getListOptions(blacklist *Blacklist) metav1.ListOptions {
	listOptions := metav1.ListOptions{}
	if blacklist != nil {
		listOptions.LabelSelector = getExcludedLabelsSelectors(blacklist)
		listOptions.FieldSelector = getExcludedFieldsSelectors(blacklist)
	}
	return listOptions
}

func getExcludedLabelsSelectors(blacklist *Blacklist) string {
	var labelsExcludedSelectors []string
	for key, value := range blacklist.Labels {
		labelsExcludedSelectors = append(labelsExcludedSelectors, fmt.Sprintf("%s!=%s", key, value))
	}
	return strings.Join(labelsExcludedSelectors, ",")
}

func getExcludedFieldsSelectors(blacklist *Blacklist) string {
	var excludedFieldSelectors []string
	for key, value := range blacklist.FieldSelectors {
		excludedFieldSelectors = append(excludedFieldSelectors, fmt.Sprintf("%s!=%s", key, value))
	}
	return strings.Join(excludedFieldSelectors, ",")
}
