package k8s

import (
	v1 "k8s.io/api/core/v1"
)

func GetDefaultContainerName(pod v1.Pod) []string {
	containerList := []string{}
	for _, container := range pod.Spec.Containers {
		containerList = append(containerList, container.Name)
	}
	for _, container := range pod.Spec.InitContainers {
		containerList = append(containerList, container.Name+" (init)")
	}
	return containerList
}
