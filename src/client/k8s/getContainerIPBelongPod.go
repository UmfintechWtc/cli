package k8s

import (
	xlogger "github.com/sirupsen/logrus"
)

func GetContainerIPBelongToAllPod(initK8sConn *K8sApiServer, namespace, podIP string) (string, []string) {
	podName := ""
	containerList := []string{}
	AllPodList := GetAllPodNames(initK8sConn, namespace)
	for _, pod := range AllPodList.Items {
		if pod.Status.PodIP == podIP {
			podName = pod.Name
			containerList = GetDefaultContainerName(pod)
			break
		}
	}
	if len(podName) == 0 {
		xlogger.Fatalf("在%s命名空间中未找到IP:[%s]归属Pod，请确认IP是否正确", namespace, podIP)
	}
	return podName, containerList
}
