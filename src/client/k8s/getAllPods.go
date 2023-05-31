package k8s

import (
	"context"

	xlogger "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetAllPodNames(initK8sConn *K8sApiServer, namespace string) *v1.PodList {
	AllPodList, err := initK8sConn.K8sClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		xlogger.Fatalf("从命名空间%s获取所有Pod数据异常: %s", namespace, err)
	}
	return AllPodList
}
