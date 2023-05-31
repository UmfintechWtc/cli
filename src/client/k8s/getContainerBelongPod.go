package k8s

import (
	"cli/src/common"
	"sort"

	xlogger "github.com/sirupsen/logrus"
)

func GetContainerBelongToAllPod(initK8sConn *K8sApiServer, namespace, containerName string) []string {
	pod_list := []string{}
	AllPodList := GetAllPodNames(initK8sConn, namespace)
	for _, pod := range AllPodList.Items {
		for _, container := range pod.Spec.Containers {
			if container.Name == containerName {
				if !common.Check_key_exists(common.Covert_Slice_To_Map(pod_list), pod.Name) {
					pod_list = append(pod_list, pod.Name)
				}
			}
		}
	}
	if len(pod_list) == 0 {
		xlogger.Fatalf("在%s命名空间中未找到容器%s归属Pod，请确认容器名称", namespace, containerName)
	}
	sort.Strings(pod_list)
	return pod_list
}
