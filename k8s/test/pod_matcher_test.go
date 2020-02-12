package k8s_test

import (
	"fmt"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	coreV1 "k8s.io/api/core/v1"
)

type ContainerMatcherConfig func(*ContainerMatcher)

type PodMatcher struct {
	containers []types.GomegaMatcher

	executed types.GomegaMatcher
}

func NewPodMatcher() *PodMatcher {
	return &PodMatcher{[]types.GomegaMatcher{}, nil}
}

func (matcher *PodMatcher) WithContainerMatching(config ContainerMatcherConfig) *PodMatcher {
	container := NewContainerMatcher()
	config(container)
	matcher.containers = append(matcher.containers, container)

	return matcher
}

func (matcher *PodMatcher) Match(actual interface{}) (bool, error) {
	pod, ok := actual.(coreV1.PodTemplateSpec)
	if !ok {
		return false, fmt.Errorf("Expected pod. Got\n%s", format.Object(actual, 1))
	}

	for _, container := range matcher.containers {
		contains := gomega.ContainElement(container)

		matcher.executed = container
		if pass, err := contains.Match(pod.Spec.Containers); !pass || err != nil {
			return pass, err
		}
	}

	return true, nil
}

func (matcher *PodMatcher) FailureMessage(actual interface{}) string {
	return matcher.executed.FailureMessage(actual)
}

func (matcher *PodMatcher) NegatedFailureMessage(actual interface{}) string {
	return matcher.executed.NegatedFailureMessage(actual)
}
